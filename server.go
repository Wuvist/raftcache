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
	node   *RaftNode
	server *grpc.Server
}

// NewGRPCHTTPServer return a server for given node
func NewGRPCHTTPServer(node *RaftNode) (server *GRPCServer, err error) {
	var n = *node

	server = &GRPCServer{
		node:   &n,
		server: grpc.NewServer(),
	}

	RegisterRaftCacheServer(server.server, server)
	return
}

// Start the grpc server, and block
func (s *GRPCServer) Start() error {
	listener, err := net.Listen("tcp", s.node.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.node.ListenAddr = listener.Addr().(*net.TCPAddr).String()
	return s.server.Serve(listener)
}

// Stop the grpc server
func (s *GRPCServer) Stop() {
	s.server.Stop()
}

func (s *GRPCServer) String() string {
	return fmt.Sprintf("Node: %s", s.node.String())
}

// Join take given node to join into group
func (s *GRPCServer) Join(ctx context.Context, in *Node) (*JoinResp, error) {
	return s.node.Join(in)
}

// Leave take given node out of group
func (s *GRPCServer) Leave(ctx context.Context, in *Node) (*LeaveResp, error) {
	return nil, nil
}
