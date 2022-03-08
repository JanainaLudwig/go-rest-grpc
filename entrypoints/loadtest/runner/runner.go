package runner

import (
	"context"
	"log"
	"sync"
	"time"
)

type Load struct {
	CallsPerSecond int
	Duration time.Duration
}

type LoadTestClient interface {
	TestFunc()
}

type Runner struct {
	ctx context.Context
	loads  []Load
	client LoadTestClient
}

func (r *Runner) AddLoad(load Load) {
	r.loads = append(r.loads, load)
}

func (r *Runner) Run() {
	for i, load := range r.loads {
		log.Printf("Running load %v - timer %v", i, load.Duration)
		//stop := make(chan bool, 1)
		//go func() {
		//	defer func() {
		//		log.Println("release")
		//		stop <- true
		//	}()
		ticker := time.NewTicker(1 * time.Second)
		timer := time.NewTimer(load.Duration)
		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()
			for {
				select {
				case <-ticker.C:
					wg.Add(load.CallsPerSecond)
					go func() {
						log.Printf("running load with %v calls", load.CallsPerSecond)
						for call := 0; call < load.CallsPerSecond; call++ {
							r.client.TestFunc()
							wg.Done()
						}
					}()
				case <- timer.C:
					log.Printf("stopping load")
					return
				}
			}

		}()

		//}()

		wg.Wait()
		//<-stop
		log.Println("finished")
	}
}