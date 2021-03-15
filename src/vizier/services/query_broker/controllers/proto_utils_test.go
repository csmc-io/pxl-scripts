package controllers_test

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	public_vizierapipb "pixielabs.ai/pixielabs/src/api/public/vizierapipb"
	"pixielabs.ai/pixielabs/src/carnot/planner/compilerpb"
	"pixielabs.ai/pixielabs/src/carnot/planner/distributedpb"
	plannerpb "pixielabs.ai/pixielabs/src/carnot/planner/plannerpb"
	"pixielabs.ai/pixielabs/src/carnot/planpb"
	"pixielabs.ai/pixielabs/src/carnot/queryresultspb"
	"pixielabs.ai/pixielabs/src/carnotpb"
	statuspb "pixielabs.ai/pixielabs/src/common/base/proto"
	typespb "pixielabs.ai/pixielabs/src/shared/types/proto"
	schemapb "pixielabs.ai/pixielabs/src/table_store/proto"
	pbutils "pixielabs.ai/pixielabs/src/utils"
	"pixielabs.ai/pixielabs/src/vizier/services/query_broker/controllers"
)

var boolScalarValuePb = `
data_type: BOOLEAN
bool_value: false
`

var queryReqPb = `
query_str: "abcd this is a test"
exec_funcs {
	func_name: "f"
	arg_values {
		name: "a"
		value: "1"
	}
	output_table_prefix: "table1"
}
exec_funcs {
	func_name: "g"
	arg_values {
		name: "c"
		value: "3"
	}
	arg_values {
		name: "d"
		value: "4"
	}
	output_table_prefix: "table2"
}
`

var executeScriptReqPb = `
query_str: "abcd this is a test"
exec_funcs {
	func_name: "f"
	arg_values {
		name: "a"
		value: "1"
	}
	output_table_prefix: "table1"
}
exec_funcs {
	func_name: "g"
	arg_values {
		name: "c"
		value: "3"
	}
	arg_values {
		name: "d"
		value: "4"
	}
	output_table_prefix: "table2"
}
`

var tablePb = `
relation {
	columns {
		column_name: "abcd"
		column_type: BOOLEAN
		column_desc: "this is a boolean column"
	}
	columns {
		column_name: "efgh"
		column_type: INT64
		column_desc: "a test column in a test table"
	}
}
name: "test"
`

var tableSemanticTypePb = `
relation {
	columns {
		column_name: "abcd"
		column_type: STRING
		column_semantic_type: ST_SERVICE_NAME
	}
	columns {
		column_name: "efgh"
		column_type: STRING
		column_semantic_type: ST_POD_NAME
	}
}
name: "test"
`

var unauthenticatedStatusPb = `
err_code: UNAUTHENTICATED
msg: "this is a message"
`

var timingStatsPb = `
timing_info {
	execution_time_ns: 10
	compilation_time_ns: 4
}
execution_stats {
	timing {
		execution_time_ns: 10
		compilation_time_ns: 4
	}
	bytes_processed: 100
	records_processed: 50
}
`

var vizierTimingStatsPb = `
execution_stats {
	timing {
		execution_time_ns: 10
		compilation_time_ns: 4
	}
	bytes_processed: 100
	records_processed: 50
}
`

var uint128Pb = `
	low: 123
	high: 456
`

var rowBatchPb = `
	eow: false
	eos: true
	num_rows: 10
	cols {
		boolean_data {
			data: true
			data: false
			data: true
		}
	}
	cols {
		string_data {
			data: "abcd"
			data: "efgh"
			data: "ijkl"
		}
	}
`

var agentPlanPb = `
dag: {
	nodes: {
		id: 1
	}
}
nodes: {
	id: 1
	dag: {
		nodes: {
			id: 2
			sorted_children: 3
		}
		nodes: {
			sorted_parents: 2
		}
	}
	nodes: {
		id: 2
		op: {
			op_type: MEMORY_SOURCE_OPERATOR
			mem_source_op: {
				name: "table1"
				column_idxs: 0
				column_idxs: 1
				column_idxs: 2
				column_names: "time_"
				column_names: "cpu_cycles"
				column_names: "upid"
				column_types: TIME64NS
				column_types: INT64
				column_types: UINT128
				tablet: "1"
			}
		}
	}
	nodes: {
		id: 3
		op: {
			op_type: GRPC_SINK_OPERATOR
			grpc_sink_op: {
				address: "foo"
				output_table {
					table_name: "agent1_table"
					column_types: TIME64NS
					column_types: INT64
					column_types: UINT128
					column_names: "time_"
					column_names: "cpu_cycles"
					column_names: "upid"
					column_semantic_types: ST_NONE
					column_semantic_types: ST_NONE
					column_semantic_types: ST_UPID
				}
			}
		}
	}
}
`

func TestVizierQueryRequestToPlannerQueryRequest(t *testing.T) {
	sv := new(public_vizierapipb.ExecuteScriptRequest)
	if err := proto.UnmarshalText(executeScriptReqPb, sv); err != nil {
		t.Fatalf("Cannot unmarshal proto")
	}

	expectedQr := new(plannerpb.QueryRequest)
	if err := proto.UnmarshalText(queryReqPb, expectedQr); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	qr, err := controllers.VizierQueryRequestToPlannerQueryRequest(sv)
	assert.Nil(t, err)
	assert.Equal(t, expectedQr, qr)
}

func TestStatusToVizierStatus(t *testing.T) {
	sv := new(statuspb.Status)
	if err := proto.UnmarshalText(unauthenticatedStatusPb, sv); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	s := controllers.StatusToVizierStatus(sv)
	assert.Equal(t, "this is a message", s.Message)
	assert.Equal(t, int32(16), s.Code)
}

func TestCompilerErrorStatusToVizierStatus(t *testing.T) {
	errs := make([]*compilerpb.CompilerError, 2)
	errs[0] = &compilerpb.CompilerError{
		Error: &compilerpb.CompilerError_LineColError{
			LineColError: &compilerpb.LineColError{
				Line:    1,
				Column:  2,
				Message: "compilation error here",
			},
		},
	}
	errs[1] = &compilerpb.CompilerError{
		Error: &compilerpb.CompilerError_LineColError{
			LineColError: &compilerpb.LineColError{
				Line:    101,
				Column:  200,
				Message: "another compilation error here",
			},
		},
	}
	compilerEG := &compilerpb.CompilerErrorGroup{
		Errors: errs,
	}
	compilerEGAny, err := types.MarshalAny(compilerEG)
	assert.Nil(t, err)
	sv := &statuspb.Status{
		Context: compilerEGAny,
	}

	s := controllers.StatusToVizierStatus(sv)
	assert.Equal(t, 2, len(s.ErrorDetails))
	assert.Equal(t, uint64(1), s.ErrorDetails[0].GetCompilerError().Line)
	assert.Equal(t, uint64(2), s.ErrorDetails[0].GetCompilerError().Column)
	assert.Equal(t, "compilation error here", s.ErrorDetails[0].GetCompilerError().Message)
	assert.Equal(t, uint64(101), s.ErrorDetails[1].GetCompilerError().Line)
	assert.Equal(t, uint64(200), s.ErrorDetails[1].GetCompilerError().Column)
	assert.Equal(t, "another compilation error here", s.ErrorDetails[1].GetCompilerError().Message)
}

func TestRelationFromTable(t *testing.T) {
	sv := new(schemapb.Table)
	if err := proto.UnmarshalText(tablePb, sv); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	expectedQm := new(public_vizierapipb.QueryMetadata)
	if err := proto.UnmarshalText(tablePb, expectedQm); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	qm, err := controllers.RelationFromTable(sv)
	assert.Nil(t, err)
	assert.Equal(t, expectedQm, qm)
}

func TestRelationFromTableWithSemanticTypes(t *testing.T) {
	sv := new(schemapb.Table)
	if err := proto.UnmarshalText(tableSemanticTypePb, sv); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	expectedQm := new(public_vizierapipb.QueryMetadata)
	if err := proto.UnmarshalText(tableSemanticTypePb, expectedQm); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	qm, err := controllers.RelationFromTable(sv)
	assert.Nil(t, err)
	assert.Equal(t, expectedQm, qm)
}

func TestQueryResultStatsToVizierStats(t *testing.T) {
	sv := new(queryresultspb.QueryResult)
	if err := proto.UnmarshalText(timingStatsPb, sv); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	expectedQd := new(public_vizierapipb.QueryData)
	if err := proto.UnmarshalText(vizierTimingStatsPb, expectedQd); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	qm := controllers.QueryResultStatsToVizierStats(sv.ExecutionStats, 4)
	assert.Equal(t, expectedQd.ExecutionStats, qm)
}

func TestUInt128ToVizierUInt128(t *testing.T) {
	sv := new(typespb.UInt128)
	if err := proto.UnmarshalText(uint128Pb, sv); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	expectedQd := new(public_vizierapipb.UInt128)
	if err := proto.UnmarshalText(uint128Pb, expectedQd); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	qm := controllers.UInt128ToVizierUInt128(sv)
	assert.Equal(t, expectedQd, qm)
}

func TestRowBatchToVizierRowBatch(t *testing.T) {
	sv := new(schemapb.RowBatchData)
	if err := proto.UnmarshalText(rowBatchPb, sv); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	expectedQd := new(public_vizierapipb.RowBatchData)
	if err := proto.UnmarshalText(rowBatchPb, expectedQd); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}

	qm, err := controllers.RowBatchToVizierRowBatch(sv, "")
	assert.Nil(t, err)
	assert.Equal(t, expectedQd, qm)
}

func TestBuildExecuteScriptResponse_RowBatch(t *testing.T) {
	receivedRB := new(schemapb.RowBatchData)
	if err := proto.UnmarshalText(rowBatchPb, receivedRB); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}
	convertedRB := new(public_vizierapipb.RowBatchData)
	if err := proto.UnmarshalText(rowBatchPb, convertedRB); err != nil {
		t.Fatalf("Cannot unmarshal proto %v", err)
	}
	convertedRB.TableID = "output_table_1_id"

	queryID := uuid.NewV4()
	queryIDpb := pbutils.ProtoFromUUID(queryID)

	msg := &carnotpb.TransferResultChunkRequest{
		Address: "foo",
		QueryID: queryIDpb,
		Result: &carnotpb.TransferResultChunkRequest_QueryResult{
			QueryResult: &carnotpb.TransferResultChunkRequest_SinkResult{
				ResultContents: &carnotpb.TransferResultChunkRequest_SinkResult_RowBatch{
					RowBatch: receivedRB,
				},
				Destination: &carnotpb.TransferResultChunkRequest_SinkResult_TableName{
					TableName: "output_table_1",
				},
			},
		},
	}
	tableIDMap := map[string]string{
		"another_table":  "another_table_id",
		"output_table_1": "output_table_1_id",
	}
	resp, err := controllers.BuildExecuteScriptResponse(msg, tableIDMap, 10)
	assert.Nil(t, err)

	assert.Nil(t, resp.Status)
	assert.Equal(t, queryID.String(), resp.QueryID)
	assert.Nil(t, resp.GetMetaData())
	assert.NotNil(t, resp.GetData())
	assert.Nil(t, resp.GetData().GetExecutionStats())
	assert.Equal(t, convertedRB, resp.GetData().GetBatch())
}

func TestBuildExecuteScriptResponse_InitiateResultStream(t *testing.T) {
	queryID := uuid.NewV4()
	queryIDpb := pbutils.ProtoFromUUID(queryID)

	msg := &carnotpb.TransferResultChunkRequest{
		Address: "foo",
		QueryID: queryIDpb,
		Result: &carnotpb.TransferResultChunkRequest_QueryResult{
			QueryResult: &carnotpb.TransferResultChunkRequest_SinkResult{
				ResultContents: &carnotpb.TransferResultChunkRequest_SinkResult_InitiateResultStream{
					InitiateResultStream: true,
				},
				Destination: &carnotpb.TransferResultChunkRequest_SinkResult_TableName{
					TableName: "output_table_1",
				},
			},
		},
	}
	tableIDMap := map[string]string{
		"another_table":  "another_table_id",
		"output_table_1": "output_table_1_id",
	}
	resp, err := controllers.BuildExecuteScriptResponse(msg, tableIDMap, 10)
	assert.Nil(t, err)
	assert.Nil(t, resp)
}

func TestBuildExecuteScriptResponse_ExecutionStats(t *testing.T) {
	queryID := uuid.NewV4()
	queryIDpb := pbutils.ProtoFromUUID(queryID)

	msg := &carnotpb.TransferResultChunkRequest{
		Address: "foo",
		QueryID: queryIDpb,
		Result: &carnotpb.TransferResultChunkRequest_ExecutionAndTimingInfo{
			ExecutionAndTimingInfo: &carnotpb.TransferResultChunkRequest_QueryExecutionAndTimingInfo{
				ExecutionStats: &queryresultspb.QueryExecutionStats{
					Timing: &queryresultspb.QueryTimingInfo{
						ExecutionTimeNs: 5010,
					},
					BytesProcessed:   4521,
					RecordsProcessed: 4,
				},
			},
		},
	}

	expectedStats := &public_vizierapipb.QueryExecutionStats{
		Timing: &public_vizierapipb.QueryTimingInfo{
			ExecutionTimeNs:   5010,
			CompilationTimeNs: 10,
		},
		BytesProcessed:   4521,
		RecordsProcessed: 4,
	}

	resp, err := controllers.BuildExecuteScriptResponse(msg, nil, 10)
	assert.Nil(t, err)

	assert.Nil(t, resp.Status)
	assert.Equal(t, queryID.String(), resp.QueryID)
	assert.Nil(t, resp.GetMetaData())
	assert.NotNil(t, resp.GetData())
	assert.Nil(t, resp.GetData().GetBatch())
	assert.Equal(t, expectedStats, resp.GetData().GetExecutionStats())
}

func TestQueryPlanResponse(t *testing.T) {
	queryIDStr := "6683eddd-0824-430c-ac0d-ce05cf9624a8"
	agentIDStr := "3ca421d4-5f85-4c99-8248-02252204e281"
	queryID, err := uuid.FromString(queryIDStr)
	if err != nil {
		t.Fatal("Error converting query ID to UUID")
	}
	agentID, err := uuid.FromString(agentIDStr)
	if err != nil {
		t.Fatal("Error converting agent ID to UUID")
	}
	planTableID := "table_plan_id"

	agentPlan := &planpb.Plan{}
	if err := proto.UnmarshalText(agentPlanPb, agentPlan); err != nil {
		t.Fatal("Cannot Unmarshal protobuf.")
	}

	planMap := make(map[uuid.UUID]*planpb.Plan)
	planMap[agentID] = agentPlan
	planMapStr := make(map[string]*planpb.Plan)
	planMapStr[agentIDStr] = agentPlan

	dagIDMap := make(map[string]uint64)
	dagIDMap[agentIDStr] = 0
	dag := &planpb.DAG{
		Nodes: []*planpb.DAG_DAGNode{
			&planpb.DAG_DAGNode{
				Id: 0,
			},
		},
	}

	plan := &distributedpb.DistributedPlan{
		QbAddressToPlan:  planMapStr,
		QbAddressToDagId: dagIDMap,
		Dag:              dag,
	}

	agentStats := []*queryresultspb.AgentExecutionStats{
		&queryresultspb.AgentExecutionStats{
			AgentID:          pbutils.ProtoFromUUID(agentID),
			ExecutionTimeNs:  123,
			BytesProcessed:   456,
			RecordsProcessed: 12,
			OperatorExecutionStats: []*queryresultspb.OperatorExecutionStats{
				&queryresultspb.OperatorExecutionStats{
					PlanFragmentId:       1,
					NodeId:               2,
					BytesOutput:          450,
					RecordsOutput:        14,
					TotalExecutionTimeNs: 50,
					SelfExecutionTimeNs:  40,
				},
				&queryresultspb.OperatorExecutionStats{
					PlanFragmentId:       1,
					NodeId:               2,
					BytesOutput:          456,
					RecordsOutput:        12,
					TotalExecutionTimeNs: 73,
					SelfExecutionTimeNs:  70,
				},
			},
		},
	}

	expected1 := []*public_vizierapipb.ExecuteScriptResponse{
		&public_vizierapipb.ExecuteScriptResponse{
			QueryID: queryIDStr,
			Result: &public_vizierapipb.ExecuteScriptResponse_Data{
				Data: &public_vizierapipb.QueryData{
					Batch: &public_vizierapipb.RowBatchData{
						TableID: "table_plan_id",
						Cols: []*public_vizierapipb.Column{
							&public_vizierapipb.Column{
								ColData: &public_vizierapipb.Column_StringData{
									StringData: &public_vizierapipb.StringColumn{
										Data: []string{
											"digraph  {\n\tsubgraph cluster_s0 {\n\t\tID = \"cluster_s0\";\n" +
												"\t\tcolor=\"lightgrey\";label=\"agent::3ca421d4-5f85-4c99-8248-02252204e281\\n" +
												"123ns\";\n\t\tn1[color=\"blue\",label=\"memory_source_operator[2]\\" +
												"nself_time: 70ns\\ntotal_time: 73ns\\nbytes: 456 B\\nrecords_processed: 12\"" +
												",shape=\"rect\"];\n\t\tn2[color=\"yellow\",label=\"grpc_sink_operator[3]\\n\"" +
												",shape=\"rect\"];\n\t\tn1->n2;\n\t\t\n\t}\n\t\n}",
										},
									},
								},
							},
						},
						NumRows: 1,
						Eow:     true,
						Eos:     true,
					},
				},
			},
		},
	}

	resp1, err := controllers.QueryPlanResponse(queryID, plan, planMap, &agentStats, planTableID, 1024*1024)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp1))
	assert.Equal(t, expected1[0], resp1[0])

	expected2 := []*public_vizierapipb.ExecuteScriptResponse{
		&public_vizierapipb.ExecuteScriptResponse{
			QueryID: queryIDStr,
			Result: &public_vizierapipb.ExecuteScriptResponse_Data{
				Data: &public_vizierapipb.QueryData{
					Batch: &public_vizierapipb.RowBatchData{
						TableID: "table_plan_id",
						Cols: []*public_vizierapipb.Column{
							&public_vizierapipb.Column{
								ColData: &public_vizierapipb.Column_StringData{
									StringData: &public_vizierapipb.StringColumn{
										Data: []string{
											"digraph  {\n\tsubgraph cluster_s0 {\n\t\tID = \"cluster_s0\";\n\t\t" +
												"color=\"lightgrey\";label=\"agent::3ca421d4-5f85-4c99-8248-02252204e281" +
												"\\n123ns\";\n\t\tn1[color=\"blue\",label=\"memory_source_operator[2]\\n" +
												"self_time: 70ns\\ntotal_time: 73ns\\nbytes: 456 B\\nrecords_processed: 12\"," +
												"shape=\"rect\"];\n\t\tn2[color=\"yellow\",label=\"grpc_sink_operator[3]\\n\"," +
												"shape=\"rect\"];\n\t\tn1->n2;",
										},
									},
								},
							},
						},
						NumRows: 1,
						Eow:     false,
						Eos:     false,
					},
				},
			},
		},
		&public_vizierapipb.ExecuteScriptResponse{
			QueryID: queryIDStr,
			Result: &public_vizierapipb.ExecuteScriptResponse_Data{
				Data: &public_vizierapipb.QueryData{
					Batch: &public_vizierapipb.RowBatchData{
						TableID: "table_plan_id",
						Cols: []*public_vizierapipb.Column{
							&public_vizierapipb.Column{
								ColData: &public_vizierapipb.Column_StringData{
									StringData: &public_vizierapipb.StringColumn{
										Data: []string{
											"\n\t\t\n\t}\n\t\n}",
										},
									},
								},
							},
						},
						NumRows: 1,
						Eow:     true,
						Eos:     true,
					},
				},
			},
		},
	}

	resp2, err := controllers.QueryPlanResponse(queryID, plan, planMap, &agentStats, planTableID, 350)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resp2))
	assert.Equal(t, expected2[0], resp2[0])
	assert.Equal(t, expected2[1], resp2[1])
}

func TestTableRelationResponses(t *testing.T) {
	queryID := uuid.NewV4()

	plannerResultPB := &distributedpb.LogicalPlannerResult{}
	if err := proto.UnmarshalText(expectedPlannerResult, plannerResultPB); err != nil {
		t.Fatal("Cannot Unmarshal protobuf.")
	}

	planPB1 := plannerResultPB.Plan.QbAddressToPlan[agent1ID]
	planPB2 := plannerResultPB.Plan.QbAddressToPlan[agent2ID]

	agentUUID1, err := uuid.FromString(agent1ID)
	if err != nil {
		t.Fatal("Error converting agent ID to UUID")
	}
	agentUUID2, err := uuid.FromString(agent2ID)
	if err != nil {
		t.Fatal("Error converting agent ID to UUID")
	}

	planMap := make(map[uuid.UUID]*planpb.Plan)
	planMap[agentUUID1] = planPB1
	planMap[agentUUID2] = planPB2

	expectedSchemaResults := make(map[string]*public_vizierapipb.ExecuteScriptResponse)
	actualSchemaResults := make(map[string]*public_vizierapipb.ExecuteScriptResponse)

	expectedSchemaResults["agent1_table"] = &public_vizierapipb.ExecuteScriptResponse{
		QueryID: queryID.String(),
		Result: &public_vizierapipb.ExecuteScriptResponse_MetaData{
			MetaData: &public_vizierapipb.QueryMetadata{
				Name: "agent1_table",
				ID:   "agent1_table_id",
				Relation: &public_vizierapipb.Relation{
					Columns: []*public_vizierapipb.Relation_ColumnInfo{
						&public_vizierapipb.Relation_ColumnInfo{ColumnName: "time_",
							ColumnType:         6,
							ColumnDesc:         "",
							ColumnSemanticType: 1,
						}, &public_vizierapipb.Relation_ColumnInfo{
							ColumnName:         "cpu_cycles",
							ColumnType:         2,
							ColumnDesc:         "",
							ColumnSemanticType: 1,
						}, &public_vizierapipb.Relation_ColumnInfo{
							ColumnName:         "upid",
							ColumnType:         3,
							ColumnDesc:         "",
							ColumnSemanticType: 200,
						},
					},
				},
			},
		},
	}
	expectedSchemaResults["agent2_table"] = &public_vizierapipb.ExecuteScriptResponse{
		QueryID: queryID.String(),
		Result: &public_vizierapipb.ExecuteScriptResponse_MetaData{
			MetaData: &public_vizierapipb.QueryMetadata{
				Name: "agent2_table",
				ID:   "agent2_table_id",
				Relation: &public_vizierapipb.Relation{
					Columns: []*public_vizierapipb.Relation_ColumnInfo{
						&public_vizierapipb.Relation_ColumnInfo{ColumnName: "time_",
							ColumnType:         6,
							ColumnDesc:         "",
							ColumnSemanticType: 1,
						},
					},
				},
			},
		},
	}

	tableIDMap := map[string]string{
		"agent1_table": "agent1_table_id",
		"agent2_table": "agent2_table_id",
	}

	resps, err := controllers.TableRelationResponses(queryID, tableIDMap, planMap)
	assert.Nil(t, err)
	assert.NotNil(t, resps)
	assert.Equal(t, 2, len(resps))
	actualSchemaResults[resps[0].GetMetaData().Name] = resps[0]
	actualSchemaResults[resps[1].GetMetaData().Name] = resps[1]

	for tableName, expected := range expectedSchemaResults {
		assert.Equal(t, expected, actualSchemaResults[tableName])
	}
}

func TestOutputSchemaFromPlan(t *testing.T) {
	agentUUIDStrs := [2]string{
		agent1ID,
		agent2ID,
	}

	agentUUIDs := make([]uuid.UUID, 0)
	for _, uid := range agentUUIDStrs {
		u, err := uuid.FromString(uid)
		if err != nil {
			t.Fatal(err)
		}
		agentUUIDs = append(agentUUIDs, u)
	}

	// Plan 1 is a valid, populated plan
	plannerResultPB := &distributedpb.LogicalPlannerResult{}
	if err := proto.UnmarshalText(expectedPlannerResult, plannerResultPB); err != nil {
		t.Fatal("Could not unmarshal protobuf text for planner result.")
	}

	planPB1 := plannerResultPB.Plan.QbAddressToPlan[agent1ID]
	// Plan 2 is an empty plan.
	planPB2 := plannerResultPB.Plan.QbAddressToPlan[agent2ID]

	planMap := make(map[uuid.UUID]*planpb.Plan)
	planMap[agentUUIDs[0]] = planPB1
	planMap[agentUUIDs[1]] = planPB2

	output := controllers.OutputSchemaFromPlan(planMap)
	assert.Equal(t, 2, len(output))
	assert.NotNil(t, output["agent1_table"])
	assert.NotNil(t, output["agent2_table"])
	assert.Equal(t, 3, len(output["agent1_table"].Columns))
	assert.Equal(t, 1, len(output["agent2_table"].Columns))

	assert.Equal(t, &schemapb.Relation_ColumnInfo{
		ColumnName:         "time_",
		ColumnType:         typespb.TIME64NS,
		ColumnSemanticType: typespb.ST_NONE,
	}, output["agent1_table"].Columns[0])

	assert.Equal(t, &schemapb.Relation_ColumnInfo{
		ColumnName:         "cpu_cycles",
		ColumnType:         typespb.INT64,
		ColumnSemanticType: typespb.ST_NONE,
	}, output["agent1_table"].Columns[1])

	assert.Equal(t, &schemapb.Relation_ColumnInfo{
		ColumnName:         "upid",
		ColumnType:         typespb.UINT128,
		ColumnSemanticType: typespb.ST_UPID,
	}, output["agent1_table"].Columns[2])

	assert.Equal(t, &schemapb.Relation_ColumnInfo{
		ColumnName:         "time_",
		ColumnType:         typespb.TIME64NS,
		ColumnSemanticType: typespb.ST_NONE,
	}, output["agent2_table"].Columns[0])
}
