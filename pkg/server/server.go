package server

import (
	pb "github.com/slntopp/rts-traffic-data-source/pkg/grpc-data-source/proto"
	"github.com/slntopp/rts-traffic-data-source/pkg/timeseries"
	"go.uber.org/zap"
)

type GrafanaDSServer struct {
	pb.UnimplementedGrafanaQueryAPIServer
	
	log *zap.Logger
	rts timeseries.TSClient
}

func NewGrafanaDSServer(log *zap.Logger, rts timeseries.TSClient) *GrafanaDSServer {
	return &GrafanaDSServer{log: log, rts: rts}
}