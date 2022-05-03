seed:
	go run entrypoints/cli/migrate.go --action seed

migrate:
	go run entrypoints/cli/migrate.go --action migrate

migrate-down:
	go run entrypoints/cli/migrate.go --action migrate:down

grpc-dev:
	docker-compose -f docker/docker-compose.yaml run grpc-dev

server:
	docker-compose -f docker/docker-compose.yaml --profile server up

server-build:
	docker-compose -f docker/docker-compose.yaml --profile server build

test-rest:
	go run entrypoints/loadtest/main.go --type rest > report-rest.csv
	
test-grpc:
	go run entrypoints/loadtest/main.go --type grpc > report-grpc.csv