package common

import (
	"context"
	"fmt"
	pb "github.com/perillaroc/takler-client/takler_protocol"
	"log"
	"time"
)

func (c *TaklerServiceClient) RunQueryShow() {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunShowRequest(ctx, &pb.ShowRequest{})

	if err != nil {
		log.Fatalf("could not init: %v", err)
	}

	fmt.Print(r.GetOutput())
}

func (c *TaklerServiceClient) RunQueryPing() {
	startTime := time.Now()
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.client.RunPingRequest(ctx, &pb.PingRequest{})

	if err != nil {
		log.Fatalf("ping server (%s:%s) failed: %v", c.Host, c.Port, err)
	}

	endTime := time.Now()
	d := endTime.Sub(startTime)

	fmt.Printf("ping server (%s:%s) succeeded in %v\n", c.Host, c.Port, d)
}
