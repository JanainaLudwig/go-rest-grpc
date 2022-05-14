package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"grpc-rest/api/handlers"
	"grpc-rest/domain"
	"log"
	"net/http"

	"github.com/bxcodec/faker/v3"
)

type Rest struct {
	ctx      context.Context
	client   http.Client
	host     string
	loadType string
}

func NewRunnerRest(host string, loadType string, loads ...Load) *Runner {
	ctx := context.Background()

	client := http.Client{}

	return &Runner{
		ctx:   ctx,
		loads: loads,
		code:  "rest",
		client: &Rest{
			client:   client,
			loadType: loadType,
			ctx:      ctx,
			host:     host,
		},
	}
}

func (r *Rest) TestFunc() error {
	if r.loadType == "post" {
		return r.postFunc()
	}

	return r.getFunc()
}

func (r *Rest) getFunc() error {
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

func (r *Rest) postFunc() error {
	values := map[string]string{"first_name": faker.FirstName(), "last_name": faker.LastName()}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequestWithContext(r.ctx, http.MethodPost, r.host, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var students handlers.ResponseId
	err = json.NewDecoder(res.Body).Decode(&students)
	if err != nil {
		return err
	}

	return nil
}
