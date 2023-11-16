package cmd

import (
	"fmt"
	"github.com/perillaroc/takler-client/common"
	"github.com/spf13/cobra"
)

/*********************************************
	requeue
 *********************************************/

type requeueCommand struct {
	BaseCommand

	host string
	port string
	//nodePaths []string
}

func newRequeueCommand() *requeueCommand {
	c := &requeueCommand{}
	requeueCmd := &cobra.Command{
		Use:   "requeue",
		Short: "[control] requeue given node(s).",
		Long:  "requeue given nodes",
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.runCommand,
	}

	requeueCmd.Flags().StringVar(&c.host, "host", "", "takler service host")
	requeueCmd.Flags().StringVar(&c.port, "port", "", "takler service port")

	c.cmd = requeueCmd
	return c
}

func (mc *requeueCommand) runCommand(cmd *cobra.Command, args []string) error {
	host, port := getHostAndPort(mc.host, mc.port)
	nodePaths := args

	fmt.Printf("%s:%s requeue: %s\n", host, port, nodePaths)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandRequeue(nodePaths)

	return nil
}

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
		Short: "[control] suspend given node(s). prevent job creation for the node and all its children nodes.",
		Long:  "suspend given nodes",
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.runCommand,
	}

	suspendCmd.Flags().StringVar(&c.host, "host", "", "takler service host")
	suspendCmd.Flags().StringVar(&c.port, "port", "", "takler service port")

	c.cmd = suspendCmd
	return c
}

func (mc *suspendCommand) runCommand(cmd *cobra.Command, args []string) error {
	host, port := getHostAndPort(mc.host, mc.port)
	nodePaths := args

	fmt.Printf("%s:%s suspend: %s\n", host, port, nodePaths)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandSuspend(nodePaths)

	return nil
}

/*********************************************
	resume
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
		Short: "[control] resume the node(s) from suspended status.",
		Long:  "resume given nodes",
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.runCommand,
	}

	resumeCmd.Flags().StringVar(&c.host, "host", "", "takler service host")
	resumeCmd.Flags().StringVar(&c.port, "port", "", "takler service port")

	c.cmd = resumeCmd
	return c
}

func (mc *resumeCommand) runCommand(cmd *cobra.Command, args []string) error {
	host, port := getHostAndPort(mc.host, mc.port)
	nodePaths := args

	fmt.Printf("%s:%s resume: %s\n", host, port, nodePaths)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandResume(nodePaths)

	return nil
}

/*********************************************
	run
 *********************************************/

type runCommand struct {
	BaseCommand

	host  string
	port  string
	force bool
	//nodePaths []string
}

func newRunCommand() *runCommand {
	c := &runCommand{}
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "[control] run the task.",
		Long:  "run the tasks, ignore triggers",
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.runCommand,
	}

	runCmd.Flags().StringVar(&c.host, "host", "", "takler service host")
	runCmd.Flags().StringVar(&c.port, "port", "", "takler service port")
	runCmd.Flags().BoolVar(&c.force, "force", false, "force run")

	c.cmd = runCmd
	return c
}

func (mc *runCommand) runCommand(cmd *cobra.Command, args []string) error {
	host, port := getHostAndPort(mc.host, mc.port)
	nodePaths := args

	fmt.Printf("%s:%s run: %s\n", host, port, nodePaths)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandRun(nodePaths, mc.force)

	return nil
}
