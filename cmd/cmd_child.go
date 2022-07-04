package cmd

import (
	"fmt"
	"github.com/perillaroc/takler-client/common"
	"github.com/spf13/cobra"
	"os"
)

type ChildCommand struct {
	BaseCommand

	childOptions struct {
		host     string
		port     string
		nodePath string
	}
}

func ignoreChangeCommand() bool {
	_, ok := os.LookupEnv("NO_TAKLER")
	if !ok {
		return false
	} else {
		return true
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
		Short: "mark task to active",
		Long:  "mark task to active",
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
	if ignoreChangeCommand() {
		fmt.Printf("ignore because NO_TAKLER is set.\n")
		return nil
	}

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
		Short: "mark task to complete",
		Long:  "mark task to complete",
		RunE:  c.runCommand,
	}

	completeCmd.Flags().StringVar(&c.childOptions.host, "host", "", "takler host")
	completeCmd.Flags().StringVar(&c.childOptions.port, "port", "", "takler port")
	completeCmd.Flags().StringVar(&c.childOptions.nodePath, "node-path", "", "node path")

	c.cmd = completeCmd
	return c
}

func (mc *completeCommand) runCommand(cmd *cobra.Command, args []string) error {
	if ignoreChangeCommand() {
		fmt.Printf("ignore because NO_TAKLER is set.\n")
		return nil
	}

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
		Short: "mark task to aborted",
		Long:  "mark task to aborted",
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
	if ignoreChangeCommand() {
		fmt.Printf("ignore because NO_TAKLER is set.\n")
		return nil
	}

	host := getHost(mc.childOptions.host)
	port := getPort(mc.childOptions.port)
	nodePath := getNodePath(mc.childOptions.nodePath)
	reason := mc.reason

	fmt.Printf("%s:%s abort %s: %s\n", host, port, nodePath, reason)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandAbort(nodePath, reason)

	return nil
}

/*********************************************
	event
 *********************************************/

type eventCommand struct {
	ChildCommand

	eventName string
}

func newEventCommand() *eventCommand {
	c := &eventCommand{}
	eventCmd := &cobra.Command{
		Use:   "event",
		Short: "change event",
		Long:  "change event",
		RunE:  c.runCommand,
	}

	eventCmd.Flags().StringVar(&c.childOptions.host, "host", "", "takler host")
	eventCmd.Flags().StringVar(&c.childOptions.port, "port", "", "takler port")
	eventCmd.Flags().StringVar(&c.childOptions.nodePath, "node-path", "", "node path")
	eventCmd.Flags().StringVar(&c.eventName, "event-name", "", "event name")
	eventCmd.MarkFlagRequired("event-name")

	c.cmd = eventCmd
	return c
}

func (mc *eventCommand) runCommand(cmd *cobra.Command, args []string) error {
	if ignoreChangeCommand() {
		fmt.Printf("ignore because NO_TAKLER is set.\n")
		return nil
	}

	host := getHost(mc.childOptions.host)
	port := getPort(mc.childOptions.port)
	nodePath := getNodePath(mc.childOptions.nodePath)
	eventName := mc.eventName

	fmt.Printf("%s:%s event %s: %s\n", host, port, nodePath, eventName)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandEvent(nodePath, eventName)

	return nil
}

/*********************************************
	meter
 *********************************************/

type meterCommand struct {
	ChildCommand

	meterName  string
	meterValue string
}

func newMeterCommand() *meterCommand {
	c := &meterCommand{}
	meterCmd := &cobra.Command{
		Use:   "meter",
		Short: "change meter",
		Long:  "change meter",
		RunE:  c.runCommand,
	}

	meterCmd.Flags().StringVar(&c.childOptions.host, "host", "", "takler host")
	meterCmd.Flags().StringVar(&c.childOptions.port, "port", "", "takler port")
	meterCmd.Flags().StringVar(&c.childOptions.nodePath, "node-path", "", "node path")
	meterCmd.Flags().StringVar(&c.meterName, "meter-name", "", "meter name")
	meterCmd.Flags().StringVar(&c.meterValue, "meter-value", "", "meter value")
	meterCmd.MarkFlagRequired("meter-name")
	meterCmd.MarkFlagRequired("meter-value")

	c.cmd = meterCmd
	return c
}

func (mc *meterCommand) runCommand(cmd *cobra.Command, args []string) error {
	if ignoreChangeCommand() {
		fmt.Printf("ignore because NO_TAKLER is set.\n")
		return nil
	}

	host := getHost(mc.childOptions.host)
	port := getPort(mc.childOptions.port)
	nodePath := getNodePath(mc.childOptions.nodePath)
	meterName := mc.meterName
	meterValue := mc.meterValue

	fmt.Printf("%s:%s meter %s: %s with %s\n", host, port, nodePath, meterName, meterValue)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunCommandMeter(nodePath, meterName, meterValue)

	return nil
}
