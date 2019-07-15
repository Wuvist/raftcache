package raftcache

//go:generate protoc raftcache.proto --go_out=plugins=grpc:.

// RaftNode represent a node in raftcache
type RaftNode struct {
	Node
}

// NewRaftNode returns new raftnode with given adminAddr, listenAddr
func NewRaftNode(adminAddr, listenAddr string) (node *RaftNode, err error) {
	node = &RaftNode{}
	node.Status = Node_ALONE
	node.AdminAddr = adminAddr
	node.ListenAddr = listenAddr
	return
}
