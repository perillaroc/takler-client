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

/*********************************************
	init
 *********************************************/
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

/*********************************************
	complete
 *********************************************/
type completeCommand struct {
	ChildCommand
}

func newCompleteCommand() *completeCommand {
	c := &completeCommand{}
	completeCmd := &cobra.Command{
		Use:   "complete",
		Short: "complete",
		Long:  "complete",
		RunE:  c.runCommand,
	}

	completeCmd.Flags().StringVar(&c.childOptions.host, "host", "", "takler host")
	completeCmd.Flags().StringVar(&c.childOptions.port, "port", "", "takler port")
	completeCmd.Flags().StringVar(&c.childOptions.nodePath, "node-path", "", "node path")

	c.cmd = completeCmd
	return c
}

func (mc *completeCommand) runCommand(cmd *cobra.Command, args []string) error {
	host := getHost(mc.childOptions.host)
	port := getPort(mc.childOptions.port)
	nodePath := getNodePath(mc.childOptions.nodePath)

	fmt.Printf("%s:%s complete %s\n", host, port, nodePath)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandComplete(nodePath)

	return nil
}

/*********************************************
	abort
 *********************************************/
type abortCommand struct {
	ChildCommand

	reason string
}

func newAbortCommand() *abortCommand {
	c := &abortCommand{}
	abortCmd := &cobra.Command{
		Use:   "abort",
		Short: "abort",
		Long:  "abort",
		RunE:  c.runCommand,
	}

	abortCmd.Flags().StringVar(&c.childOptions.host, "host", "", "takler host")
	abortCmd.Flags().StringVar(&c.childOptions.port, "port", "", "takler port")
	abortCmd.Flags().StringVar(&c.childOptions.nodePath, "node-path", "", "node path")
	abortCmd.Flags().StringVar(&c.reason, "reason", "", "abort reason")

	c.cmd = abortCmd
	return c
}

func (mc *abortCommand) runCommand(cmd *cobra.Command, args []string) error {
	host := getHost(mc.childOptions.host)
	port := getPort(mc.childOptions.port)
	nodePath := getNodePath(mc.childOptions.nodePath)
	reason := mc.reason

	fmt.Printf("%s:%s abort %s: %s\n", host, port, nodePath, reason)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandAbort(nodePath, reason)

	return nil
}
