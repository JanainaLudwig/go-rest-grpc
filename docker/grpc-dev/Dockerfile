# STEP 1
FROM golang:1.17 AS build

RUN apt-get update
RUN apt install -y protobuf-compiler


WORKDIR .

COPY /grpc ./grpc
COPY protoc.sh .


RUN go get -u google.golang.org/grpc \
    github.com/golang/protobuf/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

#RUN echo $(cat /grpc/students.proto)
#RUN ["protoc -I=/grpc/proto --go_out=/grpc/proto /grpc/students.proto"]
#RUN ["ls"]

ENTRYPOINT ["./protoc.sh"]