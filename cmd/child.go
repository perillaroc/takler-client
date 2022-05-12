package cmd

import (
	"context"
	"fmt"
	pb "github.com/perillaroc/takler-client/takler_protocol"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type ChildCommand struct {
	BaseCommand

	childOptions struct {
		host     string
		port     string
		nodePath string
	}
}

type initCommand struct {
	ChildCommand

	taskId string
}

func newInitCommand() *initCommand {
	c := &initCommand{}
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init",
		Long:  "init",
		RunE:  c.runCommand,
	}

	initCmd.Flags().StringVar(&c.childOptions.host, "host", "", "takler host")
	initCmd.Flags().StringVar(&c.childOptions.port, "port", "", "takler port")
	initCmd.Flags().StringVar(&c.childOptions.nodePath, "node-path", "", "node path")
	initCmd.Flags().StringVar(&c.taskId, "task-id", "", "task id")
	initCmd.MarkFlagRequired("task-id")

	c.cmd = initCmd
	return c
}

func (mc *initCommand) runCommand(cmd *cobra.Command, args []string) error {
	fmt.Printf("%s:%s init %s with %s\n",
		mc.childOptions.host, mc.childOptions.port, mc.childOptions.nodePath, mc.taskId)

	serverAddr := fmt.Sprintf("%s:%s", mc.childOptions.host, mc.childOptions.port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTaklerServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RunInitCommand(ctx, &pb.InitCommand{
		ChildOptions: &pb.ChildCommandOptions{
			NodePath: mc.childOptions.nodePath,
		},
		TaskId: mc.taskId,
	})

	if err != nil {
		log.Fatalf("could not init: %v", err)
	}

	log.Printf("%d", r.GetFlag())

	return nil
}
