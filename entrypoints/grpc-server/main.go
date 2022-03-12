package main

import (
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-rest/config"
	"grpc-rest/core"
	"grpc-rest/grpc/proto"
	services "grpc-rest/grpc/server"
	"log"
	"net"
)

func main() {
	config.LoadEnv(config.RootPath() + "/config/.env")

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("comparison-go-grpc"),
		newrelic.ConfigLicense(config.App.NewRelicLicence),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Fatalln(err)
	}
	config.App.SetNewrelicApp(app)

	core.StartApp(true)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(nrgrpc.UnaryServerInterceptor(app)),
		grpc.StreamInterceptor(nrgrpc.StreamServerInterceptor(app)),
	)

	proto.RegisterStudentsServiceServer(server, services.NewStudentsServiceController())

	reflection.Register(server)

	con, err := net.Listen("tcp", ":" + config.App.GrpcPort)
	if err != nil {
		log.Fatalln(err)
	}
	err = server.Serve(con)
	if err != nil {
		log.Fatalln(err)
	}
}

