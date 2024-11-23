package main

import . "bully_algorithm/worker"

func main() {
	// Initiate the workers
	workers := []Worker{}

	for i := 1; i <= 6; i++ {
		workers = append(workers, NewWorker(i))
	}

	//Introduce the workers to each other
	for i := range workers {
		peersLocs := []*Worker{}
		for j := range workers {
			if i == j {
				continue
			}

			peersLocs = append(peersLocs, &workers[j])
		}
		workers[i].AddPeers(peersLocs...)
	}

	// Turn all workers on (not sexually)
	for i := range workers {
		workers[i].Run()
	}

	workers[5].Alive = false

	workers[2].SendToAllPeers()
}
