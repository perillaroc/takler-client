package common

import (
	"context"
	"fmt"
	pb "github.com/perillaroc/takler-client/takler_protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func CreateTaklerServiceClient(host string, port string) *TaklerServiceClient {
	client := &TaklerServiceClient{
		Host: host,
		Port: port,
	}
	return client
}

type TaklerServiceClient struct {
	Host string
	Port string

	conn   *grpc.ClientConn
	client pb.TaklerServerClient
}

func (c *TaklerServiceClient) createConnection() {
	var err error
	c.conn, err = grpc.Dial(
		c.getServerAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
}

func (c *TaklerServiceClient) closeConnection() {
	c.conn.Close()
}

func (c *TaklerServiceClient) createClient() {
	c.client = pb.NewTaklerServerClient(c.conn)
}

func (c *TaklerServiceClient) getServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *TaklerServiceClient) RunCommandInit(nodePath string, taskId string) {
	c.createConnection()
	defer c.closeConnection()

	c.createClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.client.RunInitCommand(ctx, &pb.InitCommand{
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
