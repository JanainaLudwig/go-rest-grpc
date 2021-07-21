# STEP 1
FROM golang:1.16-alpine AS build

RUN apk add --no-cache git

WORKDIR /source

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

COPY . .

#RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/grpc-server entrypoints/grpc-server/main.go

# STEP 2
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build /source/out/grpc-server /app/grpc-server

COPY config/.env /source/config/.env

EXPOSE 9000

CMD ["/app/grpc-server"]
