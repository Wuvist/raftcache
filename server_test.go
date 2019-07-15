package raftcache

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestServer(t *testing.T) {
	node, err := NewRaftNode("test", ":0")
	if err != nil {
		t.Fatal(err)
	}

	server, err := NewGRPCHTTPServer(node)
	if err != nil {
		t.Fatal(err)
	}
	go server.Start()

	println(server.String())
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
	time.Sleep(10 * time.Millisecond)
	println(s1.String())

	conn, err := grpc.Dial(s1.node.ListenAddr, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := NewRaftCacheClient(conn)

	n := new(Node)
	n.Group = "test"
	resp, err := client.Join(context.Background(), n)
	if resp.Result != JoinResp_SUCCESS {
		t.Errorf("Invalid Join result %s", resp)
	}

	s1.Stop()
}
