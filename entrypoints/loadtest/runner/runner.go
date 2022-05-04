package runner

import (
	"context"
	"log"
	"sync"
	"time"
	"github.com/schollz/progressbar/v3"
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
	startTime time.Time
	endTime time.Time
}

func (r *Runner) AddLoad(load Load) {
	r.loads = append(r.loads, load)
}

func (r *Runner) Run(maxWorkers int) ReportSummary {
	if maxWorkers == 0 {
		maxWorkers = 300
	}
	runnerGroup := sync.WaitGroup{}

	totalRequests := r.getTotalRequestsCalls()
	responsesChan := make(chan RequestReport, totalRequests)

	bar := progressbar.Default(int64(totalRequests))
	r.startTime = time.Now()
	for i, load := range r.loads {
		runnerGroup.Add(1)
		loadSync := make(chan bool)
		go func(load Load, loadIndex int) {
			//log.Printf("Running load %v - duration %v", loadIndex + 1, load.Duration)
			ticker := time.NewTicker(1 * time.Second)
			timer := time.NewTimer(load.Duration)
			defer func() {
				runnerGroup.Done()
			}()
			for {
				select {
				case <-ticker.C:
					runnerGroup.Add(load.CallsPerSecond)
					go func() {
						maxGoroutines := make(chan bool, maxWorkers)

						//log.Printf("Running load with %v requests...", load.CallsPerSecond)
						for call := 0; call < load.CallsPerSecond; call++ {
							maxGoroutines <- true
							go func() {
								defer func() {
									err := bar.Add(1)
									if err != nil {
										log.Println(err)
									}
									<- maxGoroutines
									runnerGroup.Done()
								}()
								r.runCall(responsesChan)
							}()
						}
						//log.Printf("FINISHED load with %v requests...", load.CallsPerSecond)
					}()
				case <- timer.C:
					loadSync <- true
					return
				}
			}

		}(load, i)

		<- loadSync
	}

	reportControl := make(chan bool)
	go func() {
		for res := range responsesChan {
			r.report = append(r.report, res)
		}
		reportControl <- true
	}()

	runnerGroup.Wait()
	close(responsesChan)
	<-reportControl
	log.Println("Load test finished")
	r.endTime = time.Now()

	return r.GetReportSummary()
}

func (r *Runner) runCall(responses chan<- RequestReport) {
	now := time.Now()
	err := r.client.TestFunc()
	endTime := time.Now()
	responseTime := time.Since(now)

	responses <- RequestReport{
		responseTime: responseTime,
		endTime:      endTime,
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
