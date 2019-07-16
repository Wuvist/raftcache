package raftcache

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"google.golang.org/grpc"
)

var s1 *GRPCServer

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

func init() {
	node, err := NewRaftNode("test", ":0")
	if err != nil {
		panic(err)
	}

	s1, err = NewGRPCHTTPServer(node)
	if err != nil {
		panic(err)
	}
	go s1.Start()

	// Make sure s1 started
	time.Sleep(1 * time.Millisecond)
	println(s1.String())
}

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

func testJoin(t *testing.T) {
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
	s1.node.Status = Node_ALONE
}

func getClient() (RaftCacheClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(s1.node.ListenAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := NewRaftCacheClient(conn)
	return client, conn
}

func TestJoinConcurrent(t *testing.T) {
	client, conn := getClient()
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
