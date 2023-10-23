package common

import (
	"context"
	pb "github.com/perillaroc/takler-client/takler_protocol"
	"log"
	"time"
)

func (c *TaklerServiceClient) RunCommandInit(nodePath string, taskId string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunCommandInit(ctx, &pb.InitCommand{
		ChildOptions: &pb.ChildCommandOptions{
			NodePath: nodePath,
		},
		TaskId: taskId,
	})

	if err != nil {
		log.Fatalf("could not init: %v", err)
	}

	log.Printf("%d", r.GetFlag())
}

func (c *TaklerServiceClient) RunCommandComplete(nodePath string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunCommandComplete(ctx, &pb.CompleteCommand{
		ChildOptions: &pb.ChildCommandOptions{
			NodePath: nodePath,
		},
	})

	if err != nil {
		log.Fatalf("could not init: %v", err)
	}

	log.Printf("%d", r.GetFlag())
}

func (c *TaklerServiceClient) RunCommandAbort(nodePath string, reason string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunCommandAbort(ctx, &pb.AbortCommand{
		ChildOptions: &pb.ChildCommandOptions{
			NodePath: nodePath,
		},
		Reason: reason,
	})

	if err != nil {
		log.Fatalf("could not init: %v", err)
	}

	log.Printf("%d", r.GetFlag())
}

func (c *TaklerServiceClient) RunCommandEvent(nodePath string, eventName string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunCommandEvent(ctx, &pb.EventCommand{
		ChildOptions: &pb.ChildCommandOptions{
			NodePath: nodePath,
		},
		EventName: eventName,
	})

	if err != nil {
		log.Fatalf("could not init: %v", err)
	}

	log.Printf("%d", r.GetFlag())
}

func (c *TaklerServiceClient) RunCommandMeter(nodePath string, meterName string, meterValue string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunCommandMeter(ctx, &pb.MeterCommand{
		ChildOptions: &pb.ChildCommandOptions{
			NodePath: nodePath,
		},
		MeterName:  meterName,
		MeterValue: meterValue,
	})

	if err != nil {
		log.Fatalf("could not init: %v", err)
	}

	log.Printf("%d", r.GetFlag())
}
