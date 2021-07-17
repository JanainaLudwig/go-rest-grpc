# STEP 1
FROM golang:1.16-alpine AS build

RUN apk add --no-cache git

WORKDIR /source

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/go-api entrypoints/api/main.go

# STEP 2
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build /source/out/go-api /app/go-api

COPY config/.env /source/config/.env

EXPOSE 8080

CMD ["/app/go-api"]
