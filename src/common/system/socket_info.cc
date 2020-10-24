#include "src/common/system/socket_info.h"

#include <arpa/inet.h>
#include <fcntl.h>
#include <linux/inet_diag.h>
#include <linux/netlink.h>
#include <linux/rtnetlink.h>
#include <linux/sock_diag.h>
#include <linux/unix_diag.h>
#include <netinet/in.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <sys/un.h>
#include <unistd.h>

#include <cerrno>
#include <cstdio>
#include <cstring>
#include <string>
#include <utility>

#include "src/common/base/base.h"
#include "src/common/system/system.h"

namespace pl {
namespace system {

//-----------------------------------------------------------------------------
// NetlinkSocketProber
//-----------------------------------------------------------------------------

Status NetlinkSocketProber::Connect() {
  fd_ = socket(AF_NETLINK, SOCK_DGRAM, NETLINK_SOCK_DIAG);
  if (fd_ < 0) {
    return error::Internal("Could not create NETLINK_SOCK_DIAG connection. [errno=$0]", errno);
  }
  return Status::OK();
}

StatusOr<std::unique_ptr<NetlinkSocketProber>> NetlinkSocketProber::Create() {
  auto socket_prober_ptr = std::unique_ptr<NetlinkSocketProber>(new NetlinkSocketProber);
  PL_RETURN_IF_ERROR(socket_prober_ptr->Connect());
  return socket_prober_ptr;
}

StatusOr<std::unique_ptr<NetlinkSocketProber>> NetlinkSocketProber::Create(int net_ns_pid) {
  std::filesystem::path orig_net_ns_path =
      system::Config::GetInstance().proc_path() / "self/ns/net";
  int orig_net_ns_fd = open(orig_net_ns_path.string().c_str(), O_RDONLY);
  if (orig_net_ns_fd < 0) {
    return error::Internal("Could not save network namespace [path=$0]", orig_net_ns_path.string());
  }
  DEFER(close(orig_net_ns_fd););

  std::filesystem::path net_ns_path =
      system::Config::GetInstance().proc_path() / std::to_string(net_ns_pid) / "ns/net";
  int net_ns_fd = open(net_ns_path.string().c_str(), O_RDONLY);
  if (net_ns_fd < 0) {
    return error::Internal(
        "Could not create SocketProber in the network namespace of PID $0 [path=$1]", net_ns_pid,
        net_ns_path.string());
  }
  DEFER(close(net_ns_fd););

  // Switch network namespaces, so socket prober connects to the target network namespace.
  int retval = setns(net_ns_fd, 0);
  if (retval != 0) {
    return error::Internal("Could not change to network namespace of PID $0 [errno=$1]", net_ns_pid,
                           errno);
  }

  PL_ASSIGN_OR_RETURN(std::unique_ptr<NetlinkSocketProber> socket_prober_ptr, Create());

  // Switch back to original network namespace.
  ECHECK_EQ(setns(orig_net_ns_fd, 0), 0) << "Could not restore network namespace.";

  return socket_prober_ptr;
}

NetlinkSocketProber::~NetlinkSocketProber() {
  if (fd_ >= 0) {
    close(fd_);
  }
}

template <typename TDiagReqType>
Status NetlinkSocketProber::SendDiagReq(const TDiagReqType& msg_req) {
  ssize_t msg_len = sizeof(struct nlmsghdr) + sizeof(TDiagReqType);

  struct nlmsghdr msg_header = {};
  msg_header.nlmsg_len = msg_len;
  msg_header.nlmsg_type = SOCK_DIAG_BY_FAMILY;
  msg_header.nlmsg_flags = NLM_F_REQUEST | NLM_F_DUMP;

  struct iovec iov[2];
  iov[0].iov_base = &msg_header;
  iov[0].iov_len = sizeof(msg_header);
  iov[1].iov_base = const_cast<TDiagReqType*>(&msg_req);
  iov[1].iov_len = sizeof(msg_req);

  struct sockaddr_nl nl_addr = {};
  nl_addr.nl_family = AF_NETLINK;

  struct msghdr msg = {};
  msg.msg_name = &nl_addr;
  msg.msg_namelen = sizeof(nl_addr);
  msg.msg_iov = iov;
  msg.msg_iovlen = 2;

  ssize_t bytes_sent = 0;

  while (bytes_sent < msg_len) {
    ssize_t retval = sendmsg(fd_, &msg, 0);
    if (retval < 0) {
      return error::Internal("Failed to send NetLink messages [errno=$0]", errno);
    }
    bytes_sent += retval;
  }

  return Status::OK();
}

namespace {

Status ProcessDiagMsg(const struct inet_diag_msg& diag_msg, unsigned int len,
                      std::map<int, SocketInfo>* socket_info_entries) {
  if (len < NLMSG_LENGTH(sizeof(diag_msg))) {
    return error::Internal("Not enough bytes");
  }

  if (diag_msg.idiag_family != AF_INET && diag_msg.idiag_family != AF_INET6) {
    return error::Internal("Unsupported address family $0", diag_msg.idiag_family);
  }

  if (diag_msg.idiag_inode == 0) {
    // TODO(PL-1001): Investigate why inode of 0 is intermittently produced.
    // Shouldn't happen since we ask for for established connections only.
    LOG_EVERY_N(WARNING, 100)
        << "Did not expect inode of 0 for established connections...ignoring it.";
    return Status::OK();
  }

  auto iter = socket_info_entries->find(diag_msg.idiag_inode);
  ECHECK(iter == socket_info_entries->end())
      << absl::Substitute("Clobbering socket info at inode=$0", diag_msg.idiag_inode);

  SocketInfo socket_info = {};
  socket_info.family = diag_msg.idiag_family;
  socket_info.local_port = diag_msg.id.idiag_sport;
  socket_info.remote_port = diag_msg.id.idiag_dport;
  socket_info.state =
      magic_enum::enum_cast<ConnState>(diag_msg.idiag_state).value_or(ConnState::kUnknown);
  if (socket_info.family == AF_INET) {
    socket_info.local_addr = *reinterpret_cast<const struct in_addr*>(&diag_msg.id.idiag_src);
    socket_info.remote_addr = *reinterpret_cast<const struct in_addr*>(&diag_msg.id.idiag_dst);
  } else {
    socket_info.local_addr = *reinterpret_cast<const struct in6_addr*>(&diag_msg.id.idiag_src);
    socket_info.remote_addr = *reinterpret_cast<const struct in6_addr*>(&diag_msg.id.idiag_dst);
  }

  socket_info_entries->insert({diag_msg.idiag_inode, std::move(socket_info)});

  return Status::OK();
}

Status ProcessDiagMsg(const struct unix_diag_msg& diag_msg, unsigned int len,
                      std::map<int, SocketInfo>* socket_info_entries) {
  if (len < NLMSG_LENGTH(sizeof(diag_msg))) {
    return error::Internal("Not enough bytes");
  }

  if (diag_msg.udiag_family != AF_UNIX) {
    return error::Internal("Unsupported address family $0", diag_msg.udiag_family);
  }

  // Since we asked for UDIAG_SHOW_PEER in the unix_diag_req,
  // The response has additional attributes, which we parse here.
  // In particular, we are looking for the peer socket's inode number.
  const struct rtattr* attr;
  unsigned int rta_len = len - NLMSG_LENGTH(sizeof(diag_msg));
  unsigned int peer = 0;

  for (attr = reinterpret_cast<const struct rtattr*>(&diag_msg + 1); RTA_OK(attr, rta_len);
       attr = RTA_NEXT(attr, rta_len)) {
    switch (attr->rta_type) {
      case UNIX_DIAG_NAME:
        // Nothing for now, but could extract path name, if needed.
        // For this to work, one should also add UDIAG_SHOW_NAME to the request.
        break;
      case UNIX_DIAG_PEER:
        if (RTA_PAYLOAD(attr) >= sizeof(peer)) {
          peer = *reinterpret_cast<unsigned int*>(RTA_DATA(attr));
        }
    }
  }

  auto iter = socket_info_entries->find(diag_msg.udiag_ino);
  ECHECK(iter == socket_info_entries->end())
      << absl::Substitute("Clobbering socket info at inode=$0", diag_msg.udiag_ino);

  SocketInfo socket_info = {};
  socket_info.family = diag_msg.udiag_family;
  socket_info.local_port = diag_msg.udiag_ino;
  socket_info.local_addr = un_path_t{};
  socket_info.remote_port = peer;
  socket_info.remote_addr = un_path_t{};
  socket_info.state =
      magic_enum::enum_cast<ConnState>(diag_msg.udiag_state).value_or(ConnState::kUnknown);

  socket_info_entries->insert({diag_msg.udiag_ino, std::move(socket_info)});

  return Status::OK();
}

}  // namespace

template <typename TDiagMsgType>
Status NetlinkSocketProber::RecvDiagResp(std::map<int, SocketInfo>* socket_info_entries) {
  static constexpr int kBufSize = 8192;
  uint8_t buf[kBufSize];

  bool done = false;
  while (!done) {
    ssize_t num_bytes = recv(fd_, &buf, sizeof(buf), 0);
    if (num_bytes < 0) {
      return error::Internal("Receive call failed");
    }

    struct nlmsghdr* msg_header = reinterpret_cast<struct nlmsghdr*>(buf);

    for (; NLMSG_OK(msg_header, num_bytes); msg_header = NLMSG_NEXT(msg_header, num_bytes)) {
      if (msg_header->nlmsg_type == NLMSG_DONE) {
        done = true;
        break;
      }

      if (msg_header->nlmsg_type == NLMSG_ERROR) {
        return error::Internal("Netlink error");
      }

      if (msg_header->nlmsg_type != SOCK_DIAG_BY_FAMILY) {
        return error::Internal("Unexpected message type");
      }

#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wold-style-cast"
      TDiagMsgType* diag_msg = reinterpret_cast<TDiagMsgType*>(NLMSG_DATA(msg_header));
#pragma GCC diagnostic pop
      PL_RETURN_IF_ERROR(ProcessDiagMsg(*diag_msg, msg_header->nlmsg_len, socket_info_entries));
    }
  }

  return Status::OK();
}

Status NetlinkSocketProber::InetConnections(std::map<int, SocketInfo>* socket_info_entries,
                                            int conn_states) {
  struct inet_diag_req_v2 msg_req = {};
  msg_req.sdiag_protocol = IPPROTO_TCP;
  msg_req.idiag_states = conn_states;

  // Run once for IPv4.
  msg_req.sdiag_family = AF_INET;
  PL_RETURN_IF_ERROR(SendDiagReq(msg_req));
  PL_RETURN_IF_ERROR(RecvDiagResp<struct inet_diag_msg>(socket_info_entries));

  // Run again for IPv6.
  msg_req.sdiag_family = AF_INET6;
  PL_RETURN_IF_ERROR(SendDiagReq(msg_req));
  PL_RETURN_IF_ERROR(RecvDiagResp<struct inet_diag_msg>(socket_info_entries));

  return Status::OK();
}

Status NetlinkSocketProber::UnixConnections(std::map<int, SocketInfo>* socket_info_entries,
                                            int conn_states) {
  struct unix_diag_req msg_req = {};
  msg_req.sdiag_family = AF_UNIX;
  msg_req.udiag_states = conn_states;
  msg_req.udiag_show = UDIAG_SHOW_PEER;

  PL_RETURN_IF_ERROR(SendDiagReq(msg_req));
  PL_RETURN_IF_ERROR(RecvDiagResp<struct unix_diag_msg>(socket_info_entries));
  return Status::OK();
}

//-----------------------------------------------------------------------------
// PIDsByNetNamespace
//-----------------------------------------------------------------------------

StatusOr<uint32_t> NetNamespace(std::filesystem::path proc, uint32_t pid) {
  return NetNamespace(proc / std::to_string(pid));
}

StatusOr<uint32_t> NetNamespace(std::filesystem::path proc_pid) {
  std::filesystem::path net_ns_path = proc_pid / "ns/net";
  PL_ASSIGN_OR_RETURN(std::filesystem::path net_ns_link, fs::ReadSymlink(net_ns_path));
  PL_ASSIGN_OR_RETURN(uint32_t net_ns_inode_num,
                      fs::ExtractInodeNum(fs::kNetInodePrefix, net_ns_link.string()));
  return net_ns_inode_num;
}

std::map<uint32_t, std::vector<int>> PIDsByNetNamespace(std::filesystem::path proc) {
  std::map<uint32_t, std::vector<int>> result;

  for (const auto& p : std::filesystem::directory_iterator(proc)) {
    VLOG(1) << absl::Substitute("Directory: $0", p.path().string());
    int pid = 0;
    if (!absl::SimpleAtoi(p.path().filename().string(), &pid)) {
      VLOG(1) << absl::Substitute("Ignoring $0: Failed to parse pid.", p.path().string());
      continue;
    }

    StatusOr<uint32_t> net_ns_inode_num_status = NetNamespace(p);
    if (!net_ns_inode_num_status.ok()) {
      LOG(ERROR) << absl::Substitute(
          "Could not determine network namespace for pid $0. Message=$1.", pid,
          net_ns_inode_num_status.msg());
      continue;
    }

    result[net_ns_inode_num_status.ValueOrDie()].push_back(pid);
  }

  return result;
}

//-----------------------------------------------------------------------------
// SocketProberManager
//-----------------------------------------------------------------------------

NetlinkSocketProber* SocketProberManager::GetSocketProber(uint32_t ns) {
  auto iter = socket_probers_.find(ns);
  if (iter != socket_probers_.end()) {
    // Update the phase (similar to an LRU touch).
    iter->second.phase = current_phase_;

    VLOG(2) << absl::Substitute("SocketProberManager: Retrieving entry [ns=$0]", ns);
    return iter->second.socket_prober.get();
  }
  return nullptr;
}

StatusOr<NetlinkSocketProber*> SocketProberManager::CreateSocketProber(
    uint32_t ns, const std::vector<int>& pids) {
  StatusOr<std::unique_ptr<NetlinkSocketProber>> socket_prober_or = error::NotFound("");

  // Use any provided PID to create a socket into the network namespace.
  for (auto& pid : pids) {
    socket_prober_or = NetlinkSocketProber::Create(pid);
    if (socket_prober_or.ok()) {
      break;
    }
  }

  if (!socket_prober_or.ok()) {
    return error::Internal(
        "None of the provided PIDs for the provided namespace ($0) could be used to establish a "
        "netlink connection to the namespace. It is possible the namespace no longer exists. Error "
        "message for last attempt: $1",
        ns, socket_prober_or.msg());
  }

  VLOG(2) << absl::Substitute("SocketProberManager: Creating entry [ns=$0]", ns);

  // This socket prober will be in the network namespace defined by ns.
  std::unique_ptr<NetlinkSocketProber> socket_prober = socket_prober_or.ConsumeValueOrDie();
  NetlinkSocketProber* socket_prober_ptr = socket_prober.get();
  DCHECK_NE(socket_prober_ptr, nullptr);
  socket_probers_[ns] =
      TaggedSocketProber{.phase = current_phase_, .socket_prober = std::move(socket_prober)};
  return socket_prober_ptr;
}

StatusOr<NetlinkSocketProber*> SocketProberManager::GetOrCreateSocketProber(
    uint32_t ns, const std::vector<int>& pids) {
  // First check to see if an existing socket prober on the namespace exists.
  // If so, use that one.
  NetlinkSocketProber* socket_prober_ptr = GetSocketProber(ns);
  if (socket_prober_ptr != nullptr) {
    return socket_prober_ptr;
  }

  // Otherwise create a socket prober.
  return CreateSocketProber(ns, pids);
}

void SocketProberManager::Update() {
  // Toggle the phase.
  current_phase_ = current_phase_ ^ 1;

  // Remove socket probers that were not accessed in the last phase.
  auto iter = socket_probers_.begin();
  while (iter != socket_probers_.end()) {
    bool remove = (iter->second.phase == current_phase_);

    VLOG_IF(2, remove) << absl::Substitute("SocketProberManager: Removing entry [ns=$0]",
                                           iter->first);

    // Update iterator, deleting if necessary as we go.
    iter = remove ? socket_probers_.erase(iter) : ++iter;
  }
}

//-----------------------------------------------------------------------------
// SocketInfoManager
//-----------------------------------------------------------------------------

StatusOr<std::unique_ptr<SocketInfoManager>> SocketInfoManager::Create(
    std::filesystem::path proc_path, int conn_states) {
  std::unique_ptr<SocketInfoManager> socket_info_db_ptr(
      new SocketInfoManager(proc_path, conn_states));
  PL_ASSIGN_OR_RETURN(socket_info_db_ptr->socket_probers_, SocketProberManager::Create());
  return socket_info_db_ptr;
}

StatusOr<std::map<int, SocketInfo>*> SocketInfoManager::GetNamespaceConns(uint32_t pid) {
  PL_ASSIGN_OR_RETURN(uint32_t net_ns, NetNamespace(cfg_proc_path_, pid));

  // Step 1: Get the map of connections for this network namespace.
  // Create the map if it doesn't already exist.
  std::map<int, SocketInfo>* namespace_conns;

  auto ns_iter = connections_.find(net_ns);
  if (ns_iter != connections_.end()) {
    // Found a map of connections for this network namespace, so use it.
    namespace_conns = &ns_iter->second;
  } else {
    // No map of connections for this network namespace, so use a socker prober to populate one.
    PL_ASSIGN_OR_RETURN(NetlinkSocketProber * socket_prober,
                        socket_probers_->GetOrCreateSocketProber(net_ns, {static_cast<int>(pid)}));
    DCHECK(socket_prober != nullptr);

    ns_iter = connections_.insert(ns_iter, {net_ns, {}});
    namespace_conns = &ns_iter->second;

    Status s;

    s = socket_prober->InetConnections(namespace_conns, cfg_conn_states_);
    LOG_IF(ERROR, !s.ok()) << absl::Substitute("Failed to probe InetConnections [net_ns=$0 msg=$1]",
                                               net_ns, s.msg());

    s = socket_prober->UnixConnections(namespace_conns, cfg_conn_states_);
    LOG_IF(ERROR, !s.ok()) << absl::Substitute("Failed to probe UnixConnections [net_ns=$0 msg=$1]",
                                               net_ns, s.msg());

    ++num_socket_prober_calls_;
  }

  return namespace_conns;
}

StatusOr<SocketInfo*> SocketInfoManager::Lookup(uint32_t pid, uint32_t inode_num) {
  // Step 1: Get the map of connections for this network namespace.
  // Create the map if it doesn't already exist.
  std::map<int, SocketInfo>* namespace_conns;
  PL_ASSIGN_OR_RETURN(namespace_conns, GetNamespaceConns(pid));

  // Step 2: Lookup the inode.
  auto iter = namespace_conns->find(inode_num);
  if (iter == namespace_conns->end()) {
    return error::NotFound(
        "Likely not a TCP/Unix connection (might be some other socket type). Alternatively, might "
        "be looking in the wrong net namespace, which can happen if the target PID has connections "
        "in multiple namespaces.");
  }

  return &iter->second;
}

}  // namespace system
}  // namespace pl
