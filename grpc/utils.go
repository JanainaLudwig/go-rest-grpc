package grpc

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func PrototimeToTime(ptTime *timestamp.Timestamp) *time.Time {
	if ptTime == nil {
		return nil
	}

	tm := ptTime.AsTime()

	return &tm
}