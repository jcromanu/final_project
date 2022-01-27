package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	httpuserservice "github.com/jcromanu/final_project/http_service/pkg/http_user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var (
		logger      log.Logger
		userGRPC    *grpc.ClientConn
		grpcErr     error
		middlewares = []endpoint.Middleware{}
		opts        []grpc.DialOption
	)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	httpAddr := flag.String("http.addr", ":8081", "HTTP listen address")

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	addr := "localhost:8080"
	userGRPC, grpcErr = grpc.Dial(addr, opts...)
	if grpcErr != nil {
		level.Error(logger).Log("gRPC", grpcErr)
		os.Exit(-1)
	}

	repo := httpuserservice.NewHttpRespository(userGRPC, logger)
	srv := httpuserservice.NewHttpService(repo, logger)
	endpoints := httpuserservice.MakeEndpoints(srv, logger, middlewares)
	httpserver := httpuserservice.NewHTTPServer(endpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, httpserver)
	}()

	logger.Log("exit", <-errs)
}
