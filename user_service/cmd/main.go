package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
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
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	middlewares := []endpoint.Middleware{}
	grpcServerOptions := []kitGRPC.ServerOption{}

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
		level.Error(logger).Log("Error retrieviing env variables using default ones ")
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
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", envCfg.GrpcPort))

	if err != nil {
		level.Error(logger).Log("error creating listener: ", err)
		os.Exit(-1)
	}

	go func() {
		baseServer := gogrpc.NewServer(gogrpc.UnaryInterceptor(kitGRPC.Interceptor))
		pb.RegisterUserServiceServer(baseServer, userGRPCServer)
		logger.Log("transport", "HTTP", "addr", envCfg.GrpcPort)
		errs <- baseServer.Serve(grpcListener)
	}()

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
