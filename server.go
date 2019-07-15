package raftcache

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GRPCServer represent a raft cache server using gRPC
type GRPCServer struct {
	node *RaftNode
}

// NewGRPCHTTPServer return a server for given node
func NewGRPCHTTPServer(node *RaftNode) (server *GRPCServer, err error) {
	server = &GRPCServer{node}
	return
}

// Start start the server
func (s *GRPCServer) Start() error {
	listener, err := net.Listen("tcp", s.node.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.node.ListenAddr = listener.Addr().(*net.TCPAddr).String()

	grpcServer := grpc.NewServer()
	RegisterRaftCacheServer(grpcServer, s)

	go grpcServer.Serve(listener)

	return nil
}

func (s *GRPCServer) String() string {
	return fmt.Sprintf("Node: %s", s.node.String())
}

// Join take given node to join into group
func (s *GRPCServer) Join(ctx context.Context, in *Node) (*JoinResp, error) {
	return nil, nil
}

// Leave take given node out of group
func (s *GRPCServer) Leave(ctx context.Context, in *Node) (*LeaveResp, error) {
	return nil, nil
}
