package worker

import "fmt"

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
		Alive:  false,
	}
}

func (worker *Worker) Run() {
	worker.Alive = true

}

func (worker *Worker) GetId() int {
	return worker.id
}

func (worker *Worker) IsLeader() bool {
	return worker.leader
}

func (worker *Worker) BecomeLeader() {
	fmt.Printf("Worker %d is the new leader\n", worker.id)
	worker.leader = true
}

func (worker *Worker) AddPeers(peers ...*Worker) {
	worker.peers = append(worker.peers, peers...)
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

func SendToPeer(receiver *Worker) bool {
	if receiver.Alive {
		receiver.SendToAllPeers() // Make this function run asynchronously
		return true
	} else {
		return false
	}
}
