package raftcache

import (
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// GRPCServer represent a raft cache server using gRPC
type GRPCServer struct {
	node   *RaftNode
	server *grpc.Server
	wg     sync.WaitGroup
	mu     sync.Mutex
}

// NewGRPCHTTPServer return a server for given node
func NewGRPCHTTPServer(node *RaftNode) (server *GRPCServer, err error) {
	n := &RaftNode{
		Node:       node.Node,
		GroupNodes: node.GroupNodes,
	}

	server = &GRPCServer{
		node:   n,
		server: grpc.NewServer(),
	}

	RegisterRaftCacheServer(server.server, server)
	return
}

func (s *GRPCServer) getClient(listenAddr string) (client RaftCacheClient, conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(listenAddr, grpc.WithInsecure())
	if err != nil {
		return
	}

	client = NewRaftCacheClient(conn)

	return
}

// Start the grpc server, and block
func (s *GRPCServer) Start() error {
	listener, err := net.Listen("tcp", s.node.ListenAddr)
	if err != nil {
		return err
	}
	s.node.ListenAddr = listener.Addr().(*net.TCPAddr).String()
	s.wg.Done()
	return s.server.Serve(listener)
}

// Prepare add count to wait group
func (s *GRPCServer) Prepare() {
	s.wg.Add(1)
}

// Wait for binding to ListenAddr
func (s *GRPCServer) Wait() {
	s.wg.Wait()
}

// Stop the grpc server
func (s *GRPCServer) Stop() {
	s.server.Stop()
}

func (s *GRPCServer) String() string {
	return fmt.Sprintf("Node: %s", s.node.String())
}

// Ping for server's health checking
func (s *GRPCServer) Ping(ctx context.Context, in *Empty) (*Empty, error) {
	return &Empty{}, nil
}

// Join take given node to join into group
func (s *GRPCServer) Join(ctx context.Context, in *Node) (*JoinResp, error) {
	client, conn, err := s.getClient(in.ListenAddr)
	if err != nil {
		s := &JoinResp{}
		s.Result = JoinResp_PINGFAIL
		return s, nil
	}

	defer conn.Close()

	_, err = client.Ping(ctx, &Empty{})
	if err != nil {
		s := &JoinResp{}
		s.Result = JoinResp_PINGFAIL
		return s, nil
	}

	return s.node.Join(in)
}

// Leave take given node out of group
func (s *GRPCServer) Leave(ctx context.Context, in *Node) (*LeaveResp, error) {
	return s.node.Leave(in)
}
