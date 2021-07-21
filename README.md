# Go Rest / Grpc comparison

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

## Protocol buffer
### Install protoc compiler
````shell
apt install -y protobuf-compiler

go get -u google.golang.org/grpc

go get -u github.com/golang/protobuf/protoc-gen-go

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
protoc --go_out=plugins=grpc:. *.proto
````

## Poc Result
### Payload size
With 1000 registers in students database:
Rest: 149,000kB
GRPC:  45,916kB
