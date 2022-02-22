package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-kit/kit/endpoint"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
	userservice "github.com/jcromanu/final_project/details_service/pkg/details_service"
	"github.com/jcromanu/final_project/user_service/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credentials := options.Credential{
		Username: envCfg.MongoUser,
		Password: envCfg.MongoPwd,
	}
	uri := fmt.Sprintf("mongodb://%v:%v", envCfg.MongoAddr, envCfg.MongoPort)
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credentials)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		errs <- err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			errs <- err
		}
	}()

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		errs <- err
	}

	db := client.Database(envCfg.MongoDb)

	repo := userservice.NewUserRepository(db, logger)
	userService := userservice.NewService(repo, logger, validator)
	userEndpoints := userservice.MakeEndpoints(userService, logger, middlewares)
	userGRPCServer := userservice.NewGRPCServer(userEndpoints, grpcServerOptions, logger)
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", envCfg.GrpcPort))

	if err != nil {
		level.Error(logger).Log(err.Error())
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
	Hostname  string `env:"HOSTNAME" envDefault:"localhost"`
	GrpcPort  int    `env:"GRPC_SERVER_PORT" envDefault:"8080"`
	MongoUser string `env:"MONGO_USER"`
	MongoPwd  string `env:"MONGO_PASSWORD"`
	MongoDb   string `env:"MONGO_DB"`
	MongoAddr string `env:"MONGO_ADDR"`
	MongoPort string `env:"MONGO_PORT"`
}
