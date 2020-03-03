package main

import "sync"

const raftCount = 3

type Leader struct {
	Term     int
	LeaderId int
}

type Raft struct {
	mu              sync.Mutex
	me              int
	currentTerm     int
	voteFor         int
	state           int
	lastMessageTime int64
	currentLeader   int
	message         chan bool
	electCh         chan bool
	heartBeat       chan bool
	heartBearRe     chan bool
	timeout         int
}

var leader = Leader{0, -1}

func main() {

}
