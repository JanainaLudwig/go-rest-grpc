# Go Rest / Grpc comparison

## How to run
Create a **/config/.env** file. There is an example available inside the config folder.

### Docker
Download this repository and run
````shell
docker-compose -f docker/docker-compose.yaml up --build
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

## Protocol buffer
### Install protoc compiler
````shell
apt install -y protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
export PATH="$PATH:$(go env GOPATH)/bin"
````
Add this to ~/.bash_profile
```
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```

Run ``source ~/.bash_profile`` to take effect

### Compiling .proto
````shell
cd grpc
protoc --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative students.proto 
````

## Poc Result
### Payload size
With 1000 registers in students database:
Rest: 149,000kB
GRPC:  45,916kB
