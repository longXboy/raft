package raft

import (
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type StateType uint16

const (
	StateLeader StateType = iota
	StateFollower
	StateCandidate
)

type Log struct {
	term     uint64
	index    uint64
	dataType uint32
	data     interface{}
}

type Raft struct {
	// ID is the identity of the local raft. ID cannot be 0.
	ID    uint64
	peers []uint64

	currentTerm  uint64
	leader       uint64
	currentState StateType
	voteFor      uint64
	logs         []Log
	StateStore   map[string]interface{}

	commitIndex uint64
	lastApplied uint64

	nextIndex  []uint64
	matchIndex []uint64
}

func (raft *Raft) Vote(ctx context.Context, req *VoteRequest) (*VoteReply, error) {
	return &VoteReply{}, nil
}

func (raft *Raft) Append(ctx context.Context, req *AppendRequest) (*AppendReply, error) {
	return &AppendReply{}, nil
}

func NewRaft(id uint64, peers []uint64) *Raft {
	return &Raft{ID: id, peers: peers}
}

func (raft *Raft) MainLoop() {
	var lastHeartBeat time.Time
	lastTimer := time.Now()
	n := rand.Intn(150)
	timer := time.NewTimer(time.Duration(int64(150+n) * int64(time.Millisecond)))
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			if lastHeartBeat.After(lastTimer) {
				next := (time.Millisecond * 150) - time.Now().Sub(lastHeartBeat)
				timer.Reset(next)
			}
		}
	}
}

func (raft *Raft) Start() {
	go raft.MainLoop()

	lis, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterRaftRPCServer(s, raft)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
