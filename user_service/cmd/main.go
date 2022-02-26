package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/go-kit/kit/endpoint"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jcromanu/final_project/user_service/pb"
	userservice "github.com/jcromanu/final_project/user_service/pkg/user_service"
	"google.golang.org/grpc"
	gogrpc "google.golang.org/grpc"
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
	grpcServerOptions := []kitGRPC.ServerOption{}
	validator := validator.New()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Info(logger).Log("msg", "grpc service started")
	defer level.Info(logger).Log("msg", "grpc service stoped")

	envCfg := serverConfig{}
	if err := env.Parse(&envCfg); err != nil {
		level.Error(logger).Log(err.Error())
	}

	cfg := mysql.Config{
		User:   envCfg.MySQLUsr,
		Passwd: envCfg.MySQLPwd,
		Net:    envCfg.MySQLNet,
		Addr:   envCfg.MySQLAddr,
		DBName: envCfg.MYSQLDBName,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		level.Error(logger).Log(err.Error())
		os.Exit(-1)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		level.Error(logger).Log(err.Error())
		os.Exit(-1)
	}

	repo := userservice.NewUserRepository(db, logger)
	userService := userservice.NewService(repo, logger, validator)
	userEndpoints := userservice.MakeEndpoints(userService, logger, middlewares)
	userGRPCServer := userservice.NewGRPCServer(userEndpoints, grpcServerOptions, logger)
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", envCfg.GrpcPort))
	if err != nil {
		level.Error(logger).Log(err.Error())
		os.Exit(-1)
	}
	baseServer := gogrpc.NewServer(gogrpc.UnaryInterceptor(kitGRPC.Interceptor))
	pb.RegisterUserServiceServer(baseServer, userGRPCServer)
	go func() {
		logger.Log("transport", "HTTP", "addr", envCfg.GrpcPort)
		errs <- baseServer.Serve(grpcListener)
	}()
	level.Error(logger).Log("aqui no llega")

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:"+"8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		level.Error(logger).Log(err.Error())
		os.Exit(-1)
	}
	mux := runtime.NewServeMux()
	ctxCncl, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterUserServiceHandler(ctxCncl, mux, conn)
	if err != nil {
		level.Error(logger).Log(err.Error())
		os.Exit(-1)
	}

	httpLst, err := net.Listen("tcp", ":8081")
	if err != nil {
		level.Error(logger).Log(err.Error())
		os.Exit(-1)
	}

	go func() {
		logger.Log("transport", "HTTP", "addr", "8081")
		errs <- http.Serve(httpLst, mux)
	}()
	level.Error(logger).Log("aqui no llega 2")

	level.Error(logger).Log("grpc server closing on error : ", <-errs)
}

type serverConfig struct {
	Hostname    string `env:"HOSTNAME" envDefault:"localhost"`
	GrpcPort    int    `env:"GRPC_SERVER_PORT" envDefault:"8080"`
	MySQLUsr    string `env:"MY_SQL_USER"`
	MySQLPwd    string `env:"MY_SQL_PASSWORD"`
	MySQLNet    string `env:"MY_SQL_NET"`
	MySQLAddr   string `env:"MY_SQL_ADDR"`
	MYSQLDBName string `env:"MY_SQL_DB_NAME"`
}
