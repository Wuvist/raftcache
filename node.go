package raftcache

import (
	"fmt"
	"sync"
	"time"
)

//go:generate protoc raftcache.proto --go_out=plugins=grpc:.

// RaftNode represent a node in raftcache
type RaftNode struct {
	Node
	GroupNodes []Node
	muStatus   sync.Mutex
	mu         sync.Mutex
}

/*
Node Status state machine

* Once a node joined a group, it could never return Node_Alone
*/
var nodeStatueStateMachine = map[Node_Statuses][]Node_Statuses{
	Node_ALONE:        {Node_INITIATING, Node_HANDSHAKING},
	Node_INITIATING:   {Node_ALONE, Node_INGROUP},
	Node_HANDSHAKING:  {Node_INGROUP, Node_DISCONNECTED},
	Node_INGROUP:      {Node_HANDSHAKING, Node_DISCONNECTED},
	Node_DISCONNECTED: {Node_HANDSHAKING},
}

// NewRaftNode returns new raftnode with given group, listenAddr
func NewRaftNode(group, listenAddr string) (node *RaftNode, err error) {
	node = &RaftNode{}
	node.Status = Node_ALONE
	node.ListenAddr = listenAddr
	node.Group = group
	return
}

// SetStatus set the node with new status; state machine checking is enforced
func (r *RaftNode) SetStatus(status Node_Statuses) (existingStatus Node_Statuses, err error) {
	r.muStatus.Lock()
	defer r.muStatus.Unlock()

	existingStatus = r.Status
	possibleStatues := nodeStatueStateMachine[r.Status]
	for _, s := range possibleStatues {
		if s == status {
			r.Status = status
			return existingStatus, nil
		}
	}

	return existingStatus, fmt.Errorf("Not allow to set status from %s to %s", r.Status.String(), status.String())
}

func (r *RaftNode) initGroupNodes() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.GroupNodes) == 0 {
		r.GroupNodes = append(r.GroupNodes, r.Node)
	}
}

// ValidateJoin validates if given node could be joined
func (r *RaftNode) ValidateJoin(node *Node) (resp *JoinResp) {
	r.initGroupNodes()

	resp = &JoinResp{}
	resp.Result = JoinResp_SUCCESS

	if r.Group != node.Group {
		resp.Result = JoinResp_REJECTED
		resp.Message = "Invalid group name"
		return
	}

	if r.Status == Node_HANDSHAKING {
		resp.Result = JoinResp_TRYLATER
		resp.Message = "Current node is handshaking"
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

	return
}

// Join add a new node to current node
func (r *RaftNode) Join(node *Node) (resp *JoinResp, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	resp = r.ValidateJoin(node)
	if resp.Result != JoinResp_SUCCESS {
		return
	}

	var newNode Node = *node
	newNode.Status = Node_INGROUP

	if r.Status == Node_ALONE {
		r.SetStatus(Node_INGROUP)
	}

	time.Sleep(1 * time.Microsecond)
	r.GroupNodes = append(r.GroupNodes, newNode)

	return
}

// Leave take away given node from group
func (r *RaftNode) Leave(node *Node) (resp *LeaveResp, err error) {
	r.initGroupNodes()

	resp = &LeaveResp{}

	if node.ListenAddr == r.ListenAddr {
		resp.Result = LeaveResp_REJECTED
		resp.Message = "Can't leave self"
		return
	}

	if node.Group != r.Group {
		resp.Result = LeaveResp_REJECTED
		resp.Message = "Invalid group name"
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	nodes := make([]Node, 0, len(r.GroupNodes))

	for _, n := range r.GroupNodes {
		if n.ListenAddr != node.ListenAddr {
			nodes = append(nodes, n)
		}
	}

	if len(nodes) == len(r.GroupNodes) {
		resp.Result = LeaveResp_NOTINGROUP
		return
	}

	r.GroupNodes = nodes

	if len(r.GroupNodes) == 1 {
		r.SetStatus(Node_ALONE)
	}

	resp.Result = LeaveResp_SUCCESS
	return
}
