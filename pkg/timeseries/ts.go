package timeseries

import (
	"go.uber.org/zap"

	rts "github.com/RedisTimeSeries/redistimeseries-go"
)

type TSClient struct {
	log *zap.Logger
	ts *rts.Client
}

func NewTSClient(log *zap.Logger, redisHost string) TSClient {
	return TSClient{
		log: log.Named("TSClient"),
		ts: rts.NewClient(redisHost, "", nil),
	}
}