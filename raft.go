package raft

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

func NewRaft(id uint64, peers []uint64) *Raft {
	return &Raft{ID: id, peers: peers}
}
