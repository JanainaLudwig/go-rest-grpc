# STEP 1
FROM golang:1.16-alpine AS build

RUN apk add --no-cache git

WORKDIR /source

COPY go.mod .
COPY go.sum .

#RUN go mod download

COPY . .

#RUN CGO_ENABLED=0 go test -v

RUN GO_ENABLED=0 GOOS=linux go build -o ./out/loadtest entrypoints/loadtest/main.go

# STEP 2
FROM alpine:3.9
RUN apk add ca-certificates

COPY config/.env /source/config/.env
COPY --from=build /source/out/loadtest /app/loadtest
#
#COPY database/ /source/database/

#CMD ["/app/loadtest"]
