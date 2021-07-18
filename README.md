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


````go
var defaultRequestTimeout = time.Second * 10
type grpcService struct {
  grpcClient mysvcgrpc.UserServiceClient
}

func NewGRPCService(connString string) (mysvc.Service, error) {
  conn, err := grpc.Dial(connString, grpc.WithInsecure())
  if err != nil {
    return nil, err
  }
  return &grpcService{grpcClient: mysvcgrpc.NewUserServiceClient(conn)}, nil
}
func (s *grpcService) GetUsers(ids []int64) (result map[int64]mysvc.User, err error) {
  result = map[int64]mysvc.User{}
  req := &mysvcgrpc.GetUsersRequest{
    Ids: ids,
  }
  ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
  defer cancelFunc()
  resp, err := s.grpcClient.GetUsers(ctx, req)
  if err != nil {
    return
  }
  for _, grpcUser := range resp.GetUsers() {
    u := unmarshalUser(grpcUser)
    result[u.ID] = u
  }
  return
}
func (s *grpcService) GetUser(id int64) (result mysvc.User, err error) {
  req := &mysvcgrpc.GetUsersRequest{
    Ids: []int64{id},
  }
  ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
  defer cancelFunc()
  resp, err := s.grpcClient.GetUsers(ctx, req)
  if err != nil {
    return
  }
  for _, grpcUser := range resp.GetUsers() {
    // sanity check: only the requested user should be present in results
    if grpcUser.GetId() == id {
      return unmarshalUser(grpcUser), nil
    }
  }
  return result, mysvc.ErrNotFound
}
func unmarshalUser(grpcUser *mysvcgrpc.User) (result mysvc.User) {
  result.ID = grpcUser.Id
  result.Name = grpcUser.Name
  return
}
````