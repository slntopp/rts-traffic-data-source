package main

import (
	"flag"
	"net"
	"os"

	pb "github.com/slntopp/rts-traffic-data-source/pkg/grpc-data-source/proto"
	"github.com/slntopp/rts-traffic-data-source/pkg/server"
	"github.com/slntopp/rts-traffic-data-source/pkg/timeseries"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

var (
	log *zap.Logger
	redisHost string
)

func init() {
	flag.Bool("debug", false, "Drop Log Level to DEBUG(-1)")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetDefault("REDIS_HOST", "localhost:6379")

	viper.AutomaticEnv()
	
	level := 0
	if viper.GetBool("debug") {
		level = -1
	}

	atom := zap.NewAtomicLevel()
	atom.SetLevel(zapcore.Level(level))
	encoderCfg := zap.NewProductionEncoderConfig()
	log = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	redisHost = viper.GetString("REDIS_HOST")
}

func main() {
	defer func() {
		_ = log.Sync()
	}()
	log.Info("Starting Service")

	rts := timeseries.NewTSClient(log, redisHost)
	
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen", zap.String("address", ":8080"), zap.Error(err))
	}
	srv := grpc.NewServer()
	pb.RegisterGrafanaQueryAPIServer(srv, server.NewGrafanaDSServer(log, rts))

	if err := srv.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC", zap.Error(err))
	}
}