package main

import "sync"

import "math/rand"

import "time"

import "fmt"

const raftCount = 3

// Leader ...
type Leader struct {
	Term     int
	LeaderID int
}

// Raft ...
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
	for i := 0; i < raftCount; i++ {
		Make(i)
	}

	for {
	}
}

// Make ...
func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	rf.voteFor = -1
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	rf.setTerm(0)
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartBearRe = make(chan bool)
	rand.Seed(time.Now().UnixNano())
	go rf.election()
	go rf.sendLeaderHeartBeat()
	return rf
}

func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

func (rf *Raft) election() {
	var result bool
	for {
		timeout := randRange(150, 300)
		rf.lastMessageTime = millsecond()
		select {
		case <-time.After(time.Duration(timeout) * time.Microsecond):
			fmt.Println("当前节点状态：", rf.state)
		}
		result = false
		for !result {
			result = rf.electionOneRound(&leader)
		}
	}
}

func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func millsecond() int64 {
	return time.Now().Unix() / int64(time.Millisecond)
}

func (rf *Raft) electionOneRound(leader *Leader) bool {
	var timeout int64
	timeout = 100
	var vote int
	var triggerHeartbeat bool
	last := millsecond()
	success := false
	rf.mu.Lock()
	rf.becomeCandicate()
	rf.mu.Unlock()
	fmt.Println("start electing leader")
	for {
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					if leader.LeaderID < 0 {
						rf.electCh <- true
					}
				}()
			}
		}

		vote = 1
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.electCh:
				if ok {
					vote++
					success = vote > raftCount/2
					if success && !triggerHeartbeat {
						triggerHeartbeat = true
						rf.mu.Lock()
						rf.becomeLeader()
						rf.mu.Unlock()
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为leader")
						fmt.Println("leader 开始发送心跳信号")
					}
				}
			}
		}
		if timeout+last < millsecond() || (vote > raftCount/2) || rf.currentLeader > -1 {
			break
		} else {
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

func (rf *Raft) becomeCandicate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.voteFor = rf.me
	rf.currentLeader = -1
}

func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}

func (rf *Raft) sendLeaderHeartBeat() {
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}

func (rf *Raft) sendAppendEntriesImpl() {
	if rf.currentLeader == rf.me {
		var successCount = 0
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					rf.heartBearRe <- true
				}()
			}
		}
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartBearRe:
				if ok {
					successCount++
					if successCount > raftCount/2 {
						fmt.Println("投票成功，心跳信号ok")
					}
				}
			}
		}
	}
}
