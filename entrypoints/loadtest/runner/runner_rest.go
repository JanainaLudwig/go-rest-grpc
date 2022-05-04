package runner

import (
	"context"
	"encoding/json"
	"grpc-rest/domain"
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

	defer res.Body.Close()

	var students []domain.Student
	err = json.NewDecoder(res.Body).Decode(&students)
	if err != nil {
		return err
	}

	return nil
}
