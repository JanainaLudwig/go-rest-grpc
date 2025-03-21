seed:
	go run entrypoints/cli/migrate.go --action seed

migrate:
	go run entrypoints/cli/migrate.go --action migrate

migrate-down:
	go run entrypoints/cli/migrate.go --action migrate:down

grpc-dev:
	docker-compose -f docker/docker-compose.yaml run grpc-dev

server:
	docker-compose -f docker/docker-compose.yaml --profile server up -d

server-build:
	docker-compose -f docker/docker-compose.yaml --profile server build

test-rest:
	for loop in 1; do \
      go run entrypoints/loadtest/main.go --type rest --method $(method) > results/$(prefix)-rest-$$loop.csv; \
    done
	
test-grpc:
	for loop in 1; do \
      go run entrypoints/loadtest/main.go --type grpc --method $(method) > results/$(prefix)-grpc-$$loop.csv; \
    done

monitor:
	echo "Name,CPUPerc,MemUsage,NetIO,BlockIO,PIDs" > $(file).csv & \
	while true; do docker stats $(container) --no-stream --format "table {{.Name}},{{.CPUPerc}},{{.MemUsage}},{{.NetIO}},{{.BlockIO}},{{.PIDs}}" | grep 'NAME' -v | tee --append $(file).csv; sleep 0.3; done