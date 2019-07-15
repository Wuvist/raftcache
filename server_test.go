package raftcache

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestServer(t *testing.T) {
	node, err := NewRaftNode("test", "-1")
	if err != nil {
		t.Fatal(err)
	}

	server, err := NewGRPCHTTPServer(node)
	if err != nil {
		println(err)
		t.Fatal(err)
	}
	err = server.Start()
	if err == nil {
		t.Error("server should not start")
	}
}

func TestJoin(t *testing.T) {
	node, err := NewRaftNode("test", ":0")
	if err != nil {
		t.Fatal(err)
	}

	s1, err := NewGRPCHTTPServer(node)
	if err != nil {
		t.Fatal(err)
	}
	go s1.Start()

	// Make sure s1 started
	time.Sleep(1 * time.Millisecond)
	println(s1.String())

	conn, err := grpc.Dial(s1.node.ListenAddr, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := NewRaftCacheClient(conn)

	n := new(Node)
	resp, _ := client.Join(context.Background(), n)
	if resp.Result != JoinResp_REJECTED || resp.Message != "Invalid group name" {
		t.Errorf("Should not jion: %s", resp)
	}

	n.Group = "test"
	n.Status = Node_INGROUP
	resp, _ = client.Join(context.Background(), n)
	if resp.Result != JoinResp_REJECTED || resp.Message != "Invalid group status INGROUP" {
		t.Errorf("Should not jion: %s", resp)
	}

	n.Status = Node_ALONE
	n.Group = "test"
	resp, _ = client.Join(context.Background(), n)
	if resp.Result != JoinResp_SUCCESS {
		t.Errorf("Invalid Join result %s", resp)
	}

	resp, _ = client.Join(context.Background(), n)
	if resp.Result != JoinResp_ALREADYJOINED || resp.Message != "already in group" {
		t.Errorf("Invalid Join result %s", resp)
	}

	s1.node.Status = Node_DISCONNECTED
	resp, _ = client.Join(context.Background(), n)
	if resp.Result != JoinResp_REJECTED || resp.Message != "Can't join a disconnected node" {
		t.Errorf("Invalid Join result %s", resp)
	}

	s1.Stop()
}

func TestLeave(t *testing.T) {
	node, err := NewRaftNode("test", ":0")
	if err != nil {
		t.Fatal(err)
	}

	s1, err := NewGRPCHTTPServer(node)
	if err != nil {
		t.Fatal(err)
	}
	go s1.Start()

	// Make sure s1 started
	time.Sleep(1 * time.Millisecond)
	println(s1.String())

	conn, err := grpc.Dial(s1.node.ListenAddr, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := NewRaftCacheClient(conn)
	n := new(Node)
	client.Leave(context.Background(), n)
	s1.Stop()
}
