package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	httpuserservice "github.com/jcromanu/final_project/http_service/pkg/http_user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	middlewares := []endpoint.Middleware{}
	cfg := serverConfig{}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Info(logger).Log("msg", "http service started ")
	defer level.Info(logger).Log("msg", "htpp service stoped")

	if err := env.Parse(&cfg); err != nil {
		level.Error(logger).Log("Error retrieviing env variables using default ones ")
	}
	grpcAddr := fmt.Sprintf("%s:%d", cfg.Hostname, cfg.GrpcPort)
	userGRPC, grpcErr := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if grpcErr != nil {
		level.Error(logger).Log("gRPC", grpcErr)
		os.Exit(-1)
	}
	httpPort := fmt.Sprintf(":%d", cfg.HttpPort)
	httpAddr := flag.String("http.addr", httpPort, "HTTP listen address")

	repo := httpuserservice.NewRespository(userGRPC, logger)
	srv := httpuserservice.NewHttpService(repo, logger)
	endpoints := httpuserservice.MakeEndpoints(srv, logger, middlewares)
	httpserver := httpuserservice.NewHTTPServer(endpoints, logger)

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, httpserver)
	}()

	logger.Log("http client server closing on error ", <-errs)
}

type serverConfig struct {
	Hostname string `env:"HOSTNAME" envDefault:"localhost"`
	HttpPort int    `env:"HTTP_SERVER_PORT" envDefault:"8081"`
	GrpcPort int    `env:"GRPC_SERVER_PORT" envDefault:"8080"`
}
