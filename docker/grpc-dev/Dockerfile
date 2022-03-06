# STEP 1
FROM golang:1.17 AS build

RUN apt-get update
RUN apt install -y protobuf-compiler


WORKDIR .

#RUN adduser user
#RUN chown -R user:user /go/grpc
##RUN chmod -R 755 /grpc
#USER user

COPY /grpc ./grpc
#COPY /grpc/protoc.sh ./grpc/protoc.sh

RUN export PATH=$PATH:/usr/local/go/bin/go
RUN export GO111MODULE=on

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

RUN export PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /go/grpc
#RUN #sh ./protoc.sh

#ENTRYPOINT ["./protoc.sh"]