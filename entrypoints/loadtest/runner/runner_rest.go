package runner

import (
	"context"
	"log"
	"net/http"
)

type Rest struct {
	ctx context.Context
	client http.Client
	host string
}

func NewRunnerRest(host string, loads ...Load) *Runner {
	ctx := context.Background()

	client := http.Client{}

	return &Runner{
		ctx:   ctx,
		loads: loads,
		client: &Rest{
			client: client,
			ctx: ctx,
			host: host,
		},
	}
}

func (r *Rest) TestFunc()  {
	req, err := http.NewRequestWithContext(r.ctx, http.MethodGet, r.host, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = r.client.Do(req)
	if err != nil {
		log.Println(err)
	}
}
