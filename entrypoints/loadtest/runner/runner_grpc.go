package runner

import (
	"context"
	"google.golang.org/grpc"
	"grpc-rest/grpc/proto"
	"log"
	"time"
)

type RunnerGrpc struct {
	ctx context.Context
	client proto.StudentsServiceClient
	loads []Load
}

func NewRunnerGrpc(host string) *RunnerGrpc {
	ctx := context.Background()
	conn, e := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}

	client := proto.NewStudentsServiceClient(conn)

	return &RunnerGrpc{
		ctx: ctx,
		client: client,
	}
}

func (r *RunnerGrpc) AddLoad(load Load) {
	r.loads = append(r.loads, load)
}

func (r *RunnerGrpc) Run() {
	for i, load := range r.loads {
		log.Printf("Running load %v - timer %v", i, load.Duration)
		stop := make(chan bool, 1)
		go func() {
			defer func() {
				log.Println("release")
				stop <- true
			}()
			ticker := time.NewTicker(1 * time.Second)
			timer := time.NewTimer(load.Duration)

			for {
				select {
				case <-ticker.C:
					log.Printf("running load with %v calls", load.CallsPerSecond)
					for call := 0; call < load.CallsPerSecond; call++ {
						_, err := r.client.GetStudents(r.ctx, &proto.GetStudentsRequest{})
						if err != nil {
							log.Println(err)
						}
					}
				case <- timer.C:
					log.Printf("stopping load")
					return
				}
			}

		}()

		<-stop
		log.Println("finished")
	}
}