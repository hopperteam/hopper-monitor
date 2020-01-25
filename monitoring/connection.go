package monitoring

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/hopperteam/hopper-monitor/dependencies/protocol"
	"github.com/hopperteam/hopper-monitor/storage"
	"github.com/hopperteam/hopper-monitor/types"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

type Connection struct {
	configProvider *storage.ConfigProvider
	identifier string
	client pb.MonitoringClient
}

func Connect(configProvider *storage.ConfigProvider, identifier string, addr string) (*Connection, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Connection{configProvider, identifier, pb.NewMonitoringClient(conn) }, nil
}

func (connection *Connection) StreamLogs() error {
	stream, err := connection.client.StreamLogs(context.Background(), &empty.Empty{})
	if err != nil {
		return err
	}
	for {
		logEntry, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", connection.client, err)
		}

		connection.configProvider.LogStorage.StoreLogEntry( &types.LogEntry{
			Instance:  connection.identifier,
			Severity:  uint8(logEntry.Severity),
			Component: logEntry.Component,
			Timestamp: time.Unix(logEntry.Timestamp / 1000, logEntry.Timestamp % 1000),
			Message:   logEntry.Message,
		})
	}

	return nil
}
