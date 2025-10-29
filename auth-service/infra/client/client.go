package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GRPCClient struct {
	conn *grpc.ClientConn
}

func NewClient(serverAddr string) *GRPCClient {
	conn, err := grpc.NewClient(serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &GRPCClient{conn: conn}
}

func (c *GRPCClient) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *GRPCClient) Context() context.Context {
	return context.Background()
}

func (c *GRPCClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
