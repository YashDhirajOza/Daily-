package main

import (
	"fmt"
	"sync"
	"time"
)

type Fork struct {
	sync.Mutex
	id int
}

type Philosopher struct {
	id        int
	leftFork  *Fork
	rightFork *Fork
}

func (p *Philosopher) dine() {
	for {
		fmt.Printf("Philosopher %d is thinking...\n", p.id)
		time.Sleep(time.Second)
		p.rightFork.Lock()
		fmt.Printf("Philosopher %d took the right fork %d\n", p.id, p.rightFork.id)
		p.leftFork.Lock()
		fmt.Printf("Philosopher %d took the left fork %d\n", p.id, p.leftFork.id)

		fmt.Printf("Philosopher %d is eating...\n", p.id)
		time.Sleep(time.Second)

		p.leftFork.Unlock()
		fmt.Printf("Philosopher %d has put down the left fork %d\n", p.id, p.leftFork.id)

		p.rightFork.Unlock()
		fmt.Printf("Philosopher %d has put down the right fork %d\n", p.id, p.rightFork.id)
	}
}

func main() {
	numPhilosophers := 5
	forks := make([]*Fork, numPhilosophers)
	philosophers := make([]*Philosopher, numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &Fork{id: i}
	}

	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}
	for i := 0; i < numPhilosophers; i++ {
		go philosophers[i].dine()
	}
	select {}
}
