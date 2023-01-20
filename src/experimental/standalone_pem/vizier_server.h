/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#pragma once

#include <absl/container/flat_hash_map.h>
#include <grpcpp/grpcpp.h>
#include <algorithm>
#include <memory>
#include <queue>
#include <string>
#include <vector>

#include "src/api/proto/vizierpb/vizierapi.grpc.pb.h"
#include "src/api/proto/vizierpb/vizierapi.pb.h"
#include "src/carnot/carnot.h"
#include "src/carnot/engine_state.h"
#include "src/carnot/planner/compiler/compiler.h"

#include "src/carnot/planpb/plan.pb.h"
#include "src/common/base/base.h"
#include "src/common/base/statuspb/status.pb.h"
#include "src/experimental/standalone_pem/sink_server.h"
#include "src/shared/types/typespb/wrapper/types_pb_wrapper.h"
#include "src/table_store/table_store.h"

namespace px {
namespace vizier {
namespace agent {

class VizierServer final : public api::vizierpb::VizierService::Service {
 public:
  VizierServer() = delete;
  VizierServer(carnot::Carnot* carnot, px::vizier::agent::StandaloneGRPCResultSinkServer* svr,
               px::carnot::EngineState* engine_state) {
    carnot_ = carnot;
    sink_server_ = svr;
    engine_state_ = engine_state;
  }

  ::grpc::Status ExecuteScript(
      ::grpc::ServerContext*, const ::px::api::vizierpb::ExecuteScriptRequest* reader,
      ::grpc::ServerWriter<::px::api::vizierpb::ExecuteScriptResponse>* response) override {
    LOG(INFO) << "Executing Script";
    // Send schema before sending query results.
    auto compiler_state = engine_state_->CreateLocalExecutionCompilerState(0);
    auto plan = px::carnot::planner::compiler::Compiler()
                    .Compile(reader->query_str(), compiler_state.get())
                    .ConsumeValueOrDie();
    for (auto f : plan.nodes()) {
      for (auto n : f.nodes()) {
        if (n.op().op_type() == carnot::planpb::OperatorType::GRPC_SINK_OPERATOR) {
          auto output_table_info = n.op().grpc_sink_op();
          if (!output_table_info.has_output_table()) {
            continue;
          }
          ::px::api::vizierpb::ExecuteScriptResponse schema_resp;
          auto metadata = schema_resp.mutable_meta_data();
          metadata->set_name(output_table_info.output_table().table_name());
          metadata->set_id(output_table_info.output_table().table_name());
          auto rel = metadata->mutable_relation();
          for (int i = 0; i < output_table_info.output_table().column_names().size(); i++) {
            auto col = rel->add_columns();
            col->set_column_name(output_table_info.output_table().column_names()[i]);
            switch (output_table_info.output_table().column_types()[i]) {
              case types::BOOLEAN:
                col->set_column_type(px::api::vizierpb::BOOLEAN);
                break;
              case types::INT64:
                col->set_column_type(px::api::vizierpb::INT64);
                break;
              case types::UINT128:
                col->set_column_type(px::api::vizierpb::UINT128);
                break;
              case types::FLOAT64:
                col->set_column_type(px::api::vizierpb::FLOAT64);
                break;
              case types::STRING:
                col->set_column_type(px::api::vizierpb::STRING);
                break;
              case types::TIME64NS:
                col->set_column_type(px::api::vizierpb::TIME64NS);
                break;
              default:
                break;
            }
          }
          response->Write(schema_resp);
        }
      }
    }

    auto query_id = sole::uuid4();

    sink_server_->AddConsumer(query_id, response);
    auto s = carnot_->ExecuteQuery(reader->query_str(), query_id, 0);
    if (s != Status::OK()) {
      return ::grpc::Status::CANCELLED;
    }

    return ::grpc::Status::OK;
  }

  ::grpc::Status HealthCheck(
      ::grpc::ServerContext*, const ::px::api::vizierpb::HealthCheckRequest*,
      ::grpc::ServerWriter<::px::api::vizierpb::HealthCheckResponse>*) override {
    return ::grpc::Status::OK;
  }

  ::grpc::Status GenerateOTelScript(::grpc::ServerContext*,
                                    const ::px::api::vizierpb::GenerateOTelScriptRequest*,
                                    ::px::api::vizierpb::GenerateOTelScriptResponse*) override {
    return ::grpc::Status(grpc::StatusCode::UNIMPLEMENTED, "Not yet implemented");
  }

 protected:
  carnot::Carnot* carnot_;
  px::vizier::agent::StandaloneGRPCResultSinkServer* sink_server_;
  px::carnot::EngineState* engine_state_;
};

class VizierGRPCServer {
 public:
  VizierGRPCServer() = delete;
  VizierGRPCServer(int port, carnot::Carnot* carnot,
                   px::vizier::agent::StandaloneGRPCResultSinkServer* svr,
                   carnot::EngineState* engine_state)
      : vizier_server_(std::make_unique<VizierServer>(carnot, svr, engine_state)) {
    grpc::ServerBuilder builder;

    std::string uri = absl::Substitute("0.0.0.0:$0", port);
    builder.AddListeningPort(uri, grpc::InsecureServerCredentials());
    builder.RegisterService(vizier_server_.get());

    grpc_server_ = builder.BuildAndStart();
    CHECK(grpc_server_ != nullptr);

    LOG(INFO) << "Starting Vizier service: " << uri;
  }

  void Stop() {
    if (grpc_server_) {
      grpc_server_->Shutdown();
    }
    grpc_server_.reset(nullptr);
  }

  ~VizierGRPCServer() { Stop(); }

  std::unique_ptr<api::vizierpb::VizierService::StubInterface> StubGenerator(
      const std::string&) const {
    grpc::ChannelArguments args;
    return px::api::vizierpb::VizierService::NewStub(grpc_server_->InProcessChannel(args));
  }

 private:
  std::unique_ptr<grpc::Server> grpc_server_;
  std::unique_ptr<VizierServer> vizier_server_;
};

}  // namespace agent
}  // namespace vizier
}  // namespace px