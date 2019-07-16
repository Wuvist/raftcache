package raftcache

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"google.golang.org/grpc"
)

// An AtomicInt is an int64 to be accessed atomically.
// Take from https://github.com/golang/groupcache/blob/master/groupcache.go#L467
type AtomicInt int64

// Add atomically adds n to i.
func (i *AtomicInt) Add(n int64) {
	atomic.AddInt64((*int64)(i), n)
}

// Get atomically gets the value of i.
func (i *AtomicInt) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}

func (i *AtomicInt) String() string {
	return strconv.FormatInt(i.Get(), 10)
}

func getServer(group, listenAddr string) *GRPCServer {
	node, err := NewRaftNode(group, listenAddr)
	if err != nil {
		panic(err)
	}

	s, err := NewGRPCHTTPServer(node)
	if err != nil {
		panic(err)
	}
	s.Prepare()
	go s.Start()

	s.Wait()
	return s
}

func TestServerFail(t *testing.T) {
	node, err := NewRaftNode("test", "-1")
	if err != nil {
		t.Fatal(err)
	}

	server, err := NewGRPCHTTPServer(node)
	if err != nil {
		t.Fatal(err)
	}
	err = server.Start()
	if err == nil {
		t.Error("server should not start")
	}

	println(server.String())
}

func TestJoin(t *testing.T) {
	s := getServer("test", ":0")
	client, conn := getClient(s)
	defer conn.Close()

	n := new(Node)
	resp, err := client.Join(context.Background(), n)
	if err != nil {
		t.Error(err)
	}
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

	s.node.Status = Node_DISCONNECTED
	resp, _ = client.Join(context.Background(), n)
	if resp.Result != JoinResp_REJECTED || resp.Message != "Can't join a disconnected node" {
		t.Errorf("Invalid Join result %s", resp)
	}

	s.Stop()
}

func getClient(s *GRPCServer) (RaftCacheClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(s.node.ListenAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := NewRaftCacheClient(conn)
	return client, conn
}

func TestJoinConcurrent(t *testing.T) {
	s := getServer("test", ":0")
	client, conn := getClient(s)
	defer conn.Close()

	var wg sync.WaitGroup

	n := new(Node)
	n.Status = Node_ALONE
	n.Group = "test"
	var successCount AtomicInt

	join := func() {
		resp, _ := client.Join(context.Background(), n)
		if resp.Result == JoinResp_SUCCESS {
			successCount.Add(1)
		}

		wg.Done()
	}

	wg.Add(3)

	go join()
	go join()
	go join()

	wg.Wait()

	if successCount.Get() != 1 {
		t.Errorf("Wrong number of join success: %d", successCount.Get())
	}
}

func TestLeave(t *testing.T) {
	s := getServer("test", ":0")
	client, conn := getClient(s)
	defer conn.Close()

	n := new(Node)
	client.Leave(context.Background(), n)
	s.Stop()
}

func TestPing(t *testing.T) {
	s := getServer("test", ":0")
	defer s.Stop()

	client, conn := getClient(s)
	defer conn.Close()

	n := new(Empty)
	_, err := client.Ping(context.Background(), n)
	if err != nil {
		t.Error(err)
	}
}
