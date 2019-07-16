package raftcache

import (
	"sync"
	"time"
)

//go:generate protoc raftcache.proto --go_out=plugins=grpc:.

// RaftNode represent a node in raftcache
type RaftNode struct {
	Node
	GroupNodes []Node
}

var mu sync.Mutex

// NewRaftNode returns new raftnode with given group, listenAddr
func NewRaftNode(group, listenAddr string) (node *RaftNode, err error) {
	node = &RaftNode{}
	node.Status = Node_ALONE
	node.ListenAddr = listenAddr
	node.Group = group
	return
}

// Join add a new node to current node
func (r *RaftNode) Join(node *Node) (resp *JoinResp, err error) {
	mu.Lock()
	defer mu.Unlock()

	resp = &JoinResp{}
	if r.Group != node.Group {
		resp.Result = JoinResp_REJECTED
		resp.Message = "Invalid group name"
		return
	}

	if r.Status == Node_DISCONNECTED {
		resp.Result = JoinResp_REJECTED
		resp.Message = "Can't join a disconnected node"
		return
	}

	if node.Status == Node_INGROUP {
		resp.Result = JoinResp_REJECTED
		resp.Message = "Invalid group status INGROUP"
		return
	}

	if len(r.GroupNodes) == 0 {
		r.GroupNodes = append(r.GroupNodes, r.Node)
	}

	for _, n := range r.GroupNodes {
		if n.ListenAddr == node.ListenAddr {
			resp.Result = JoinResp_ALREADYJOINED
			if n.ListenAddr == r.ListenAddr {
				resp.Message = "Can't join self"
			} else {
				resp.Message = "Already in group"
			}

			return
		}
	}

	var newNode Node = *node
	newNode.Status = Node_INGROUP

	if r.Status == Node_ALONE {
		r.Status = Node_INGROUP
	}

	time.Sleep(1 * time.Microsecond)
	r.GroupNodes = append(r.GroupNodes, newNode)

	return
}
