package raftcache

import "testing"

func TestServer(t *testing.T) {
	node, err := NewRaftNode(":0", ":0")
	if err != nil {
		t.Fatal(err)
	}

	server, err := NewGRPCHTTPServer(node)
	if err != nil {
		t.Fatal(err)
	}
	server.Start()

	println(server.String())
}
