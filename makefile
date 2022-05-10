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
	for loop in 1 2 3 4 5; do \
      go run entrypoints/loadtest/main.go --type rest --method $(method) > report-rest.csv && mv report-rest.csv results/$(prefix)-rest-$$loop.csv; \
    done
	
test-grpc:
	for loop in 1 2 3 4 5; do \
      go run entrypoints/loadtest/main.go --type grpc --method $(method) > report-grpc.csv && mv report-grpc.csv results/$(prefix)-grpc-$$loop.csv; \
    done
