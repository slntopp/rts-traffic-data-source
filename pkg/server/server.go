package server

import (
	pb "github.com/slntopp/rts-traffic-data-source/pkg/proto"
	"github.com/slntopp/rts-traffic-data-source/pkg/timeseries"
	"go.uber.org/zap"
)

type GrafanaDSServer struct {
	pb.UnimplementedGrafanaJSONDataSourceServiceServer
	
	log *zap.Logger
	rts timeseries.TSClient
}

func NewGrafanaDSServer(log *zap.Logger, rts timeseries.TSClient) *GrafanaDSServer {
	return &GrafanaDSServer{log: log, rts: rts}
}