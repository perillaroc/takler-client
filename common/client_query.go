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
