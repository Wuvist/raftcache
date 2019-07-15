package raftcache

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GRPCHTTPServer represent a raft cache server using gRPC for caching, and http for admin
type GRPCHTTPServer struct {
	node *RaftNode
}

// NewGRPCHTTPServer return a server for given node
func NewGRPCHTTPServer(node *RaftNode) (server *GRPCHTTPServer, err error) {
	server = &GRPCHTTPServer{node}
	return
}

// Start start the server
func (s *GRPCHTTPServer) Start() error {
	listener, err := net.Listen("tcp", s.node.AdminAddr)
	if err != nil {
		panic(err)
	}

	s.node.AdminAddr = listener.Addr().(*net.TCPAddr).String()

	lis, err := net.Listen("tcp", s.node.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.node.ListenAddr = lis.Addr().(*net.TCPAddr).String()

	grpcServer := grpc.NewServer()
	RegisterRaftCacheServer(grpcServer, s)

	go grpcServer.Serve(lis)

	return nil
}

func (s *GRPCHTTPServer) String() string {
	return fmt.Sprintf("Node: %s", s.node.String())
}

// Join take given node to join into group
func (s *GRPCHTTPServer) Join(ctx context.Context, in *Node) (*JoinResp, error) {
	return nil, nil
}

// Leave take given node out of group
func (s *GRPCHTTPServer) Leave(ctx context.Context, in *Node) (*LeaveResp, error) {
	return nil, nil
}
