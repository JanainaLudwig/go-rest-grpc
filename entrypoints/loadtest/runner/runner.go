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
	TestFunc() error
}

type Runner struct {
	ctx context.Context
	loads  []Load
	client LoadTestClient
	report []RequestReport
	code string
}

func (r *Runner) AddLoad(load Load) {
	r.loads = append(r.loads, load)
}

func (r *Runner) Run() ReportSummary {
	wg := sync.WaitGroup{}

	totalRequests := r.getTotalRequestsCalls()
	responsesChan := make(chan RequestReport, totalRequests)

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
							r.runCall(responsesChan)
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

	reportControl := make(chan bool)
	go func() {
		for res := range responsesChan {
			r.report = append(r.report, res)
		}
		reportControl <- true
	}()

	wg.Wait()
	close(responsesChan)
	<-reportControl
	log.Println("Load test finished")

	return r.GetReportSummary()
}

func (r *Runner) runCall(responses chan<- RequestReport) {
	now := time.Now()
	err := r.client.TestFunc()
	responseTime := time.Since(now)

	responses <- RequestReport{
		responseTime: responseTime,
		success:      err == nil,
		error:        err,
	}
}

func (r *Runner) getTotalRequestsCalls() int {
	totalRequests := 0
	for _, load := range r.loads {
		totalRequests += load.CallsPerSecond * int(load.Duration.Seconds())
	}

	return totalRequests
}
