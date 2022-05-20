package cmd

import (
	"fmt"
	"github.com/perillaroc/takler-client/common"
	"github.com/spf13/cobra"
)

/*********************************************
	suspend
 *********************************************/

type suspendCommand struct {
	BaseCommand

	host string
	port string
	//nodePaths []string
}

func newSuspendCommand() *suspendCommand {
	c := &suspendCommand{}
	suspendCmd := &cobra.Command{
		Use:   "suspend",
		Short: "suspend",
		Long:  "suspend",
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.runCommand,
	}

	suspendCmd.Flags().StringVar(&c.host, "host", "", "takler host")
	suspendCmd.Flags().StringVar(&c.port, "port", "", "takler port")

	c.cmd = suspendCmd
	return c
}

func (mc *suspendCommand) runCommand(cmd *cobra.Command, args []string) error {
	host := getHost(mc.host)
	port := getPort(mc.port)
	nodePaths := args

	fmt.Printf("%s:%s resume: %s\n", host, port, nodePaths)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandSuspend(nodePaths)

	return nil
}

/*********************************************
	suspend
 *********************************************/

type resumeCommand struct {
	BaseCommand

	host string
	port string
	//nodePaths []string
}

func newResumeCommand() *resumeCommand {
	c := &resumeCommand{}
	resumeCmd := &cobra.Command{
		Use:   "resume",
		Short: "resume",
		Long:  "resume",
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.runCommand,
	}

	resumeCmd.Flags().StringVar(&c.host, "host", "", "takler host")
	resumeCmd.Flags().StringVar(&c.port, "port", "", "takler port")

	c.cmd = resumeCmd
	return c
}

func (mc *resumeCommand) runCommand(cmd *cobra.Command, args []string) error {
	host := getHost(mc.host)
	port := getPort(mc.port)
	nodePaths := args

	fmt.Printf("%s:%s resume: %s\n", host, port, nodePaths)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandResume(nodePaths)

	return nil
}
