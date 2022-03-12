# Go Rest / Grpc comparison

## How to run
Create a **/config/.env** file. There is an example available inside the config folder.

### Docker
Download this repository and run
````shell
docker-compose -f docker/docker-compose.yaml up --build
````

#### Profiles
Profiles can be used to start only a group of containers, and not all of them. The available profiles are:
- ***db***: only db containers
- ***server***: only db and server containers

**Example**

``docker-compose -f docker/docker-compose.yaml --profile db up``

### With local go installation

Install golang, the run
````shell
go mod vendor
go mod download
go run entrypoints/api/main.go
````

### Protocol buffer
There is a configured container for compiling the protocol buffers. Run:

````shell
docker-compose -f docker/docker-compose.yaml run grpc-dev
sh protoc.sh
````

## Migrations
Migration are run on the startup of the applications. To create a new migration, run:
```shell
go run entrypoints/cli/migrate.go --action create --name create_subjects_table
```
This will create two files in database/migrations folder (up and down file).
To run the migrations without starting the application, run `` go run entrypoints/cli/migrate.go --action migrate``

## Poc Result (WIP)
### Payload size
With 1000 registers in students database:
Rest: 149,000kB
GRPC:  45,916kB
