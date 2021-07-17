# Go Api Template

## How to run
Create a **/config/.env** file. There is an example available inside the config folder.

### Docker
Download this repository and run
````shell
docker-compose -f docker\docker-compose.yaml up --build
````

### With local go installation
If you don't have docker installed, you can run with a local go installation.

#### Install golang

#### Run
````shell
go mod vendor
go mod download
go run entrypoints/api/main.go
````


```shell
// install protoc compiler
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```