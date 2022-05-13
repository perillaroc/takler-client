package cmd

import (
	"fmt"
	"github.com/perillaroc/takler-client/common"
	"github.com/spf13/cobra"
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
	host := getHost(mc.childOptions.host)
	port := getPort(mc.childOptions.port)
	nodePath := getNodePath(mc.childOptions.nodePath)

	taskId := mc.taskId
	fmt.Printf("%s:%s init %s with %s\n", host, port, nodePath, taskId)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandInit(nodePath, taskId)

	return nil
}
