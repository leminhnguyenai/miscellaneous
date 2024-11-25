package main

import (
	. "bully_algorithm/worker"
	"fmt"
	"sync"
)

func main() {
	// Initiate the workers
	workers := []Worker{}

	for i := 1; i <= 14; i++ {
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
	var wg sync.WaitGroup
	run := Init()

	for i := 0; i < len(workers); i++ {
		// time.Sleep(time.Millisecond * 100)
		wg.Add(1)
		go func() {
			defer wg.Done()
			run(&workers[i])
		}()
	}

	wg.Wait()

	fmt.Println()
	fmt.Println("-- ELECTION RESULTS --")
	for i := range workers {
		if workers[i].IsLeader() {
			fmt.Printf("Worker %d is the leader 👑\n", i+1)
		} else {
			fmt.Printf("Worker %d is a worker\n", i+1)
		}
	}
}
