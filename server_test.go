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

	if server.String() != `Node: Group:"test" ListenAddr:"-1" ` {
		t.Error("server should not start")
	}
}

func TestJoinSelf(t *testing.T) {
	s1 := getServer("test", ":0")
	defer s1.Stop()

	resp, _ := s1.peerJoin(s1.node.ListenAddr)

	if resp.Result != JoinResp_ALREADYJOINED || resp.Message != "Can't join self" {
		t.Errorf("Invalid Join result %s", resp)
	}
}

func TestJoin(t *testing.T) {
	s1 := getServer("test", ":0")
	defer s1.Stop()

	s2 := getServer("test", ":0")
	defer s2.Stop()

	s1.node.Group = ""
	resp, _ := s1.peerJoin(s2.node.ListenAddr)
	if resp.Result != JoinResp_REJECTED || resp.Message != "Invalid group name" {
		t.Errorf("Should not jion: %s", resp)
	}

	s1.node.Group = "test"
	s1.node.Status = Node_INGROUP
	_, err := s1.peerJoin(s2.node.ListenAddr)
	if err == nil || err.Error() != "Not allow to set status from INGROUP to INITIATING" {
		t.Errorf("Should not jion: %s", resp)
	}

	s1.node.Status = Node_ALONE
	s1.node.Group = "test"
	resp, _ = s1.peerJoin(s2.node.ListenAddr)
	if resp.Result != JoinResp_SUCCESS {
		t.Errorf("Invalid Join result %s", resp)
	}

	s1.node.Status = Node_ALONE
	resp, _ = s1.peerJoin(s2.node.ListenAddr)
	if resp.Result != JoinResp_ALREADYJOINED || resp.Message != "Already in group" {
		t.Errorf("Invalid Join result %s", resp)
	}

	s2.node.Status = Node_DISCONNECTED
	resp, err = s1.peerJoin(s2.node.ListenAddr)
	if resp.Result != JoinResp_REJECTED || resp.Message != "Can't join a disconnected node" {
		t.Errorf("Invalid Join result %s", resp)
	}
}

func checkServerGroup(t *testing.T, serverGroup []*GRPCServer) {
	for _, testServer := range serverGroup {
		for _, checkServer := range serverGroup {
			if !checkServerInGroup(testServer.node.ListenAddr, checkServer) {
				t.Errorf("Cann't find server %s in group", testServer.node.ListenAddr)
			}
		}
	}
}

func checkServerInGroup(listenAddr string, s *GRPCServer) bool {
	for _, n := range s.node.GroupNodes {
		if listenAddr == n.ListenAddr {
			return true
		}
	}

	return false
}

func TestMultiJoin(t *testing.T) {
	s1 := getServer("test", ":0")
	defer s1.Stop()

	s2 := getServer("test", ":0")
	defer s2.Stop()

	s3 := getServer("test", ":0")
	defer s2.Stop()

	resp, _ := s1.peerJoin(s2.node.ListenAddr)
	if resp.Result != JoinResp_SUCCESS {
		t.Errorf("Invalid Join result %s", resp)
	}

	checkServerGroup(t, []*GRPCServer{s1, s2})

	resp, _ = s3.peerJoin(s2.node.ListenAddr)
	if resp.Result != JoinResp_SUCCESS {
		t.Errorf("Invalid Join result %s", resp)
	}

	checkServerGroup(t, []*GRPCServer{s1, s2, s3})
}

func TestMultiJoinConcurrnt(t *testing.T) {
	s1 := getServer("test", ":0")
	defer s1.Stop()

	s2 := getServer("test", ":0")
	defer s2.Stop()

	s3 := getServer("test", ":0")
	defer s2.Stop()

	c := make(chan *JoinResp)

	j := func(s *GRPCServer) {
		resp, _ := s.peerJoin(s1.node.ListenAddr)
		c <- resp
	}

	go j(s2)
	go j(s3)

	r1, r2 := <-c, <-c
	if r1.Result == JoinResp_SUCCESS {
		if r2.Result != JoinResp_TRYLATER {
			t.Errorf("Join error: %s %s", r1.Result.String(), r2.Result.String())
		}
	}

	if r2.Result == JoinResp_SUCCESS {
		if r1.Result != JoinResp_TRYLATER {
			t.Errorf("Join error: %s %s", r1.Result.String(), r2.Result.String())
		}
	}

	if r1.Result != JoinResp_SUCCESS && r2.Result != JoinResp_SUCCESS {
		t.Errorf("Join error: %s %s", r1.Result.String(), r2.Result.String())
	}
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
	s1 := getServer("test", ":0")
	defer s1.Stop()

	client, conn := getClient(s1)
	defer conn.Close()

	s2 := getServer("test", ":0")
	defer s2.Stop()

	var wg sync.WaitGroup

	n := new(Node)
	n.ListenAddr = s2.node.ListenAddr
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
	defer s.Stop()
	client, conn := getClient(s)
	defer conn.Close()

	n := new(Node)
	resp, err := client.Leave(context.Background(), n)
	if err != nil {
		t.Error(err)
	}

	if resp.Result != LeaveResp_REJECTED || resp.Message != "Invalid group name" {
		t.Error("Shouldn't leave: " + resp.String())
	}

	n.Group = "test"
	resp, err = client.Leave(context.Background(), n)
	if err != nil {
		t.Error(err)
	}

	if resp.Result != LeaveResp_NOTINGROUP || resp.Message != "" {
		t.Error("Shouldn't leave: " + resp.String())
	}

	n.ListenAddr = s.node.ListenAddr
	resp, err = client.Leave(context.Background(), n)
	if err != nil {
		t.Error(err)
	}

	if resp.Result != LeaveResp_REJECTED || resp.Message != "Can't leave self" {
		t.Error("Shouldn't leave: " + resp.String())
	}

	s2 := getServer("test", ":0")
	defer s2.Stop()
	n.ListenAddr = s2.node.ListenAddr
	n.Status = Node_ALONE
	client.Join(context.Background(), n)

	resp, err = client.Leave(context.Background(), n)
	if err != nil {
		t.Error(err)
	}

	if resp.Result != LeaveResp_SUCCESS {
		t.Error("Fail to leave: " + resp.String())
	}
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
