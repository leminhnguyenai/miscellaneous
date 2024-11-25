package worker

import (
	"fmt"
)

type Worker struct {
	// properties with lower case letter is private (what the actually fk)
	id     int
	peers  []*Worker
	leader bool
	Alive  bool
}

func NewWorker(id int) Worker {
	return Worker{
		id:     id,
		peers:  []*Worker{},
		leader: false,
		Alive:  true,
	}
}

// Return a function that take a point worker as an agruement
func Init() func(worker *Worker) {
	lock := false
	return func(worker *Worker) {
		if lock {
			return
		}
		lock = true
		if worker.IsLeader() || worker.HasLeader() || !worker.Alive {
			return
		}
		worker.SendToAllPeers()
		lock = false
	}

}

func (sender *Worker) SendToAllPeers() {
	for i := range sender.peers {
		if sender.peers[i].GetId() < sender.GetId() {
			continue
		}
		res := SendToPeer(sender.peers[i])
		if res {
			return
		}
	}
	sender.BecomeLeader()
}

func (worker *Worker) GetId() int {
	return worker.id
}

func (worker *Worker) IsLeader() bool {
	return worker.leader
}

func (worker *Worker) HasLeader() bool {
	for i := range worker.peers {
		if worker.peers[i].IsLeader() {
			return true
		}
	}

	return false
}

func (worker *Worker) BecomeLeader() {
	if worker.HasLeader() {
		fmt.Printf("Worker %d can't be leader, there is already a leader", worker.GetId())
		return
	}
	fmt.Printf("Worker %d is the new leader \n", worker.id)
	worker.leader = true
}

func (worker *Worker) AddPeers(peers ...*Worker) {
	worker.peers = append(worker.peers, peers...)
}

func SendToPeer(receiver *Worker) bool {
	if receiver.Alive {
		receiver.SendToAllPeers() // Make this function run asynchronously
		return true
	} else {
		return false
	}
}
