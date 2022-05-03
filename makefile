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
