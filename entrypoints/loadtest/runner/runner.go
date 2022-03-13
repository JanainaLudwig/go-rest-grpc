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
	wg := sync.WaitGroup{}
	for i, load := range r.loads {
		wg.Add(1)
		loadSync := make(chan bool)
		go func(load Load, loadIndex int) {
			log.Printf("Running load %v - timer %v", loadIndex, load.Duration)
			ticker := time.NewTicker(1 * time.Second)
			timer := time.NewTimer(load.Duration)
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
					log.Printf("stopping load %v", loadIndex)
					loadSync <- true
					return
				}
			}

		}(load, i)

		<- loadSync
		log.Println("finished")
	}

	wg.Wait()
	log.Println("Load test finished")
}