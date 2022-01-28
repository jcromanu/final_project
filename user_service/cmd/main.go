package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/endpoint"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-sql-driver/mysql"
	"github.com/jcromanu/final_project/user_service/pb"
	userservice "github.com/jcromanu/final_project/user_service/pkg/user_service"
	gogrpc "google.golang.org/grpc"
)

func main() {

	var logger log.Logger

	var (
		middlewares       = []endpoint.Middleware{}
		grpcServerOptions = []kitGRPC.ServerOption{}
	)

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger, "grpc_server", "time", log.DefaultTimestamp, "caller")

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service stoped")

	//Change user and password for env variable and dockerize it
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost",
		DBName: "user_db",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		level.Error(logger).Log("mysql connection error: ", err)
		os.Exit(-1)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		level.Error(logger).Log("mysql ping error: ", err)
		os.Exit(-1)
	}

	repo := userservice.NewUserRepository(db, logger)
	userService := userservice.NewService(repo, logger)
	userEndpoints := userservice.MakeEndpoints(userService, logger, middlewares)
	userGRPCServer := userservice.NewGRPCServer(userEndpoints, grpcServerOptions, logger)
	grpcListener, err := net.Listen("tcp", ":8080")

	if err != nil {
		level.Error(logger).Log("error creating listener: ", err)
		os.Exit(-1)
	}

	baseServer := gogrpc.NewServer(gogrpc.UnaryInterceptor(kitGRPC.Interceptor))
	pb.RegisterUserServiceServer(baseServer, userGRPCServer)
	if err := baseServer.Serve(grpcListener); err != nil {
		level.Error(logger).Log("error serving grpc server", err)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {}()

	level.Error(logger).Log("exit: ", <-errs)
}
