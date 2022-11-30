// Implement the dining philosopher’s problem with the following constraints/modifications.
// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// The host allows no more than 2 philosophers to eat concurrently.
// Each philosopher is numbered, 1 through 5.
// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

package main

import (
	"fmt"
	"os"
	"sync"
)

const MAX_DINING = 2
const DINNING_GOAL = 3

type Chopstick struct {
	mu sync.Mutex
}

type Philosopher struct {
	number       int
	LeftCS       *Chopstick
	RightCS      *Chopstick
	DinningCount int
	mu           sync.Mutex
}

type Host struct {
	CurrentlyDinning int
	mu               sync.Mutex
}

func Init() ([]*Philosopher, *Host) {
	var chopsticks []Chopstick
	var philosophers []*Philosopher
	host := Host{CurrentlyDinning: 0}

	for i := 0; i < 5; i++ {
		chopsticks = append(chopsticks, Chopstick{})
	}

	philosophers = append(philosophers, &Philosopher{number: 1, LeftCS: &chopsticks[0], RightCS: &chopsticks[1]})
	philosophers = append(philosophers, &Philosopher{number: 2, LeftCS: &chopsticks[1], RightCS: &chopsticks[2]})
	philosophers = append(philosophers, &Philosopher{number: 3, LeftCS: &chopsticks[2], RightCS: &chopsticks[3]})
	philosophers = append(philosophers, &Philosopher{number: 4, LeftCS: &chopsticks[3], RightCS: &chopsticks[4]})
	philosophers = append(philosophers, &Philosopher{number: 5, LeftCS: &chopsticks[4], RightCS: &chopsticks[0]})

	// fmt.Println("Created philosophers: ", philosophers)
	return philosophers, &host
}

func (h *Host) AskForPermission(permission chan bool) {
	// fmt.Println("asking again")
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.CurrentlyDinning < MAX_DINING {
		permission <- true
	}

}

func (h *Host) IncreaseDinning() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.CurrentlyDinning >= MAX_DINING {
		fmt.Println("Error: Trying to increase over max dining capacity")
		os.Exit(1)
	}
	// fmt.Println("INcreasing ", h.CurrentlyDinning)
	h.CurrentlyDinning++
}

func (h *Host) DecreaseDinning() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.CurrentlyDinning == 0 {
		return
	}
	// fmt.Println("DEcreasing", h.CurrentlyDinning)
	h.CurrentlyDinning--
}

func (p *Philosopher) UseChopsticks() {
	p.LeftCS.mu.Lock()
	p.RightCS.mu.Lock()
	p.LeftCS.mu.Unlock()
	p.RightCS.mu.Unlock()
}

func (p *Philosopher) IncreaseDinning() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.DinningCount += 1
}
func (p *Philosopher) GetDinningCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.DinningCount
}

func (p *Philosopher) Dinning(c1 chan bool) {
	fmt.Println("starting to eat ", p.number)
	p.UseChopsticks()
	fmt.Println("finishing eating ", p.number)

	c1 <- true
}

func main() {
	philosophers, host := Init()
	c1 := make(chan bool)

	for _, p := range philosophers {
		permission := make(chan bool)

		go host.AskForPermission(permission)

		for {
			<-permission
			host.IncreaseDinning()
			p.IncreaseDinning()
			go p.Dinning(c1)
			host.DecreaseDinning()

			if p.GetDinningCount() < DINNING_GOAL {
				go host.AskForPermission(permission)
			}
			if p.GetDinningCount() == DINNING_GOAL {
				break
			}
		}

	}

	for range philosophers {
		<-c1
		<-c1
		<-c1
	}

	for _, p := range philosophers {
		fmt.Printf("philosopher %v dinned %v times\n", p.number, p.DinningCount)
	}

}
