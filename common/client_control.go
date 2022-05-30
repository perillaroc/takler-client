package common

import (
	"context"
	pb "github.com/perillaroc/takler-client/takler_protocol"
	"log"
	"time"
)

func (c *TaklerServiceClient) RunCommandSuspend(nodePaths []string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunSuspendCommand(ctx, &pb.SuspendCommand{
		NodePath: nodePaths,
	})

	if err != nil {
		log.Fatalf("could not suspend: %v", err)
	}

	log.Printf("%d", r.GetFlag())
}

func (c *TaklerServiceClient) RunCommandResume(nodePaths []string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunResumeCommand(ctx, &pb.SuspendCommand{
		NodePath: nodePaths,
	})

	if err != nil {
		log.Fatalf("could not resume: %v", err)
	}

	log.Printf("%d", r.GetFlag())
}
