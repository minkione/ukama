package client

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/ukama/ukama/systems/subscriber/hlr/pb/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Hlr struct {
	conn    *grpc.ClientConn
	timeout time.Duration
	client  pb.HlrServiceClient
	host    string
}

func NewHlr(host string, timeout time.Duration) *Hlr {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	client := pb.NewHlrServiceClient(conn)

	return &Hlr{
		conn:    conn,
		client:  client,
		timeout: timeout,
		host:    host,
	}
}

func NewHlrFromClient(hlrClient pb.HlrServiceClient) *Hlr {
	return &Hlr{
		host:    "localhost",
		timeout: 1 * time.Second,
		conn:    nil,
		client:  hlrClient,
	}
}

func (r *Hlr) Close() {
	r.conn.Close()
}

func (h *Hlr) ReadRecord(req *pb.GetRecordReq) (*pb.GetRecordResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	return h.client.Get(ctx, req)
}

func (h *Hlr) UpdateGuti(req *pb.UpdateGutiReq) (*pb.UpdateGutiResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	return h.client.UpdateGuti(ctx, req)
}

func (h *Hlr) UpdateTai(req *pb.UpdateTaiReq) (*pb.UpdateTaiResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	return h.client.UpdateTai(ctx, req)
}
