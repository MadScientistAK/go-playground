package main

import (
	"fmt"
	"sync"
	"time"
)

type chopS struct {
	sync.Mutex
}

type philosopher struct {
	id, count        int
	lCStick, rCStick *chopS
}

func (p philosopher) eat(c chan *philosopher, eg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p
		if p.count < 3 {
			p.lCStick.Lock()
			p.rCStick.Lock()
			fmt.Printf("\nPhilosopher%d has started eating", p.id)
			p.count++
			fmt.Printf("\nPhilosopher%d has finished eating", p.id)
			p.lCStick.Unlock()
			p.rCStick.Unlock()
			eg.Done()
		}
	}
}

func host(c chan *philosopher, eg *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			// Some Delay
			time.Sleep(time.Second)
		}
	}
}

func main() {
	count := 5
	var eg sync.WaitGroup
	c := make(chan *philosopher, 2)
	eg.Add(15)

	chopSticks := make([]*chopS, count)
	for i := 0; i < count; i++ {
		chopSticks[i] = new(chopS)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = new(philosopher)
		philosophers[i].id = i + 1
		philosophers[i].count = 0
		philosophers[i].lCStick = chopSticks[i]
		philosophers[i].rCStick = chopSticks[(i+1)%count]
	}

	go host(c, &eg)
	for i := 0; i < count; i++ {
		go philosophers[i].eat(c, &eg)
	}
	eg.Wait()
}
