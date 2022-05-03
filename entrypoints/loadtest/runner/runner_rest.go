package runner

import (
	"context"
	"io/ioutil"
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
		code: "rest",
		client: &Rest{
			client: client,
			ctx: ctx,
			host: host,
		},
	}
}

func (r *Rest) TestFunc() error {
	req, err := http.NewRequestWithContext(r.ctx, http.MethodGet, r.host, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return err
	}
	ioutil.ReadAll(res.Body)
	res.Body.Close()

	return err
}
