package raftcache

//go:generate protoc raftcache.proto --go_out=plugins=grpc:.

// RaftNode represent a node in raftcache
type RaftNode struct {
	Node
}

// NewRaftNode returns new raftnode with given group, listenAddr
func NewRaftNode(group, listenAddr string) (node *RaftNode, err error) {
	node = &RaftNode{}
	node.Status = Node_ALONE
	node.ListenAddr = listenAddr
	node.Group = group
	return
}
