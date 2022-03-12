seed:
	go run entrypoints/cli/migrate.go --action seed

migrate:
	go run entrypoints/cli/migrate.go --action migrate

migrate-down:
	go run entrypoints/cli/migrate.go --action migrate:down
