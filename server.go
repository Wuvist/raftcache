package raftcache

import (
	"context"
	"errors"
	"fmt"
	"log"
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

// peerJoin join the server with peer
func (s *GRPCServer) peerJoin(listenAddr string) (*JoinResp, error) {
	_, err := s.node.SetStatus(Node_INITIATING, "")
	if err != nil {
		return nil, err
	}

	client, conn, err := s.getClient(listenAddr)
	if err != nil {
		s.node.SetStatus(Node_ALONE, "")
		return nil, err
	}
	defer conn.Close()

	result, err := client.Join(context.Background(), &s.node.Node)
	if err != nil {
		s.node.SetStatus(Node_ALONE, "")
		return result, err
	}

	if result.Result != JoinResp_SUCCESS {
		s.node.SetStatus(Node_ALONE, "")
		return result, errors.New("Join server failed: " + result.Result.String())
	}

	s.node.SetStatus(Node_INGROUP, "")
	groupNodes := make([]Node, len(result.GroupNodes))
	for i, n := range result.GroupNodes {
		groupNodes[i] = *n
	}

	s.node.GroupNodes = groupNodes

	return result, nil
}

func (s *GRPCServer) peerPing(listenAddr string) error {
	client, conn, err := s.getClient(listenAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = client.Ping(context.Background(), &Empty{})
	if err != nil {
		return err
	}

	return nil
}

func (s *GRPCServer) peerHandshake(listenAddr string, in *Node) (*HandshakeResp, error) {
	println("peerHandshake: " + listenAddr)
	client, conn, err := s.getClient(listenAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return client.Handshake(context.Background(), in)
}

func (s *GRPCServer) peerHandshakeConfirm(listenAddr string, in *Node) (*HandshakeConfirmResp, error) {
	println("peerHandshakeConfirm: " + listenAddr)
	client, conn, err := s.getClient(listenAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return client.HandshakeConfirm(context.Background(), in)
}

// Join take given node to join into group
func (s *GRPCServer) Join(ctx context.Context, in *Node) (*JoinResp, error) {
	r := s.node.ValidateJoin(in)
	if r.Result != JoinResp_SUCCESS {
		return r, nil
	}

	existingStatus, err := s.node.SetStatus(Node_HANDSHAKING, in.ListenAddr)
	if err != nil {
		r.Result = JoinResp_TRYLATER
		r.Message = err.Error()
		return r, nil
	}

	err = s.peerPing(in.ListenAddr)
	if err != nil {
		s.node.SetStatus(existingStatus, "")
		r.Result = JoinResp_PINGFAIL
		return r, nil
	}

	for _, node := range s.node.GroupNodes {
		if node.ListenAddr == s.node.ListenAddr {
			continue
		}
		handshakeResult, err := s.peerHandshake(node.ListenAddr, in)
		if err != nil || handshakeResult.Result != HandshakeResp_SUCCESS {
			s.node.SetStatus(existingStatus, "")
			r.Result = JoinResp_PINGFAIL
			return r, nil
		}
	}

	for _, node := range s.node.GroupNodes {
		if node.ListenAddr == s.node.ListenAddr {
			continue
		}
		handshakeConfirmResult, err := s.peerHandshakeConfirm(node.ListenAddr, in)
		if err != nil || handshakeConfirmResult.Result != HandshakeConfirmResp_SUCCESS {
			log.Println("ConfirmJoin error: " + node.ListenAddr + " " + handshakeConfirmResult.String())
		}
	}

	s.node.SetStatus(Node_INGROUP, "")

	return s.node.Join(in)
}

// HandshakeConfirm confirm join requet for peer nodes
func (s *GRPCServer) HandshakeConfirm(ctx context.Context, in *Node) (*HandshakeConfirmResp, error) {
	r := &HandshakeConfirmResp{}

	joinResult, err := s.node.Join(in)
	if err != nil {
		r.Result = HandshakeConfirmResp_FAILED
		r.Message = err.Error()
		return r, nil
	}

	if joinResult.Result != JoinResp_SUCCESS {
		r.Result = HandshakeConfirmResp_FAILED
		r.Message = joinResult.Result.String()
		return r, nil
	}

	s.node.SetStatus(Node_INGROUP, "")
	r.Result = HandshakeConfirmResp_SUCCESS

	return r, nil
}

// Handshake forwards join request to peer node for handshake validation
func (s *GRPCServer) Handshake(ctx context.Context, in *Node) (*HandshakeResp, error) {
	r := &HandshakeResp{}
	_, err := s.node.SetStatus(Node_HANDSHAKING, in.ListenAddr)
	if err != nil {
		r.Result = HandshakeResp_TRYLATER
		r.Message = "Status wrong, try later"
		return r, nil
	}

	s.node.handshakingNode = in.ListenAddr

	err = s.peerPing(in.ListenAddr)
	if err != nil {
		s.node.SetStatus(Node_INGROUP, "")
		r.Result = HandshakeResp_PINGFAIL
		return r, nil
	}

	// todo: now must set timeout
	r.Result = HandshakeResp_SUCCESS

	return r, nil
}

// Leave take given node out of group
func (s *GRPCServer) Leave(ctx context.Context, in *Node) (*LeaveResp, error) {
	return s.node.Leave(in)
}
