package cmd

import (
	"fmt"
	"github.com/perillaroc/takler-client/common"
	"github.com/spf13/cobra"
)

/*********************************************
	show
 *********************************************/
type showCommand struct {
	BaseCommand

	host string
	port string

	showParameter bool
	showTrigger   bool
	showLimit     bool
	showEvent     bool
	showMeter     bool
	showAll       bool
}

func newShowCommand() *showCommand {
	c := &showCommand{}
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "[query] print bunch tree.",
		Long:  "print state of all flows in server",
		RunE:  c.runCommand,
	}

	showCmd.Flags().StringVar(&c.host, "host", "", "takler host")
	showCmd.Flags().StringVar(&c.port, "port", "", "takler port")
	showCmd.Flags().BoolVar(&c.showTrigger, "show-trigger", false, "show trigger")
	showCmd.Flags().BoolVar(&c.showParameter, "show-parameter", false, "show parameters")
	showCmd.Flags().BoolVar(&c.showLimit, "show-limit", true, "show limits")
	showCmd.Flags().BoolVar(&c.showEvent, "show-event", true, "show events")
	showCmd.Flags().BoolVar(&c.showMeter, "show-meter", true, "show meters")
	showCmd.Flags().BoolVar(&c.showAll, "show-all", false, "show all items, ignore other show options")

	c.cmd = showCmd
	return c
}

func (mc *showCommand) runCommand(cmd *cobra.Command, args []string) error {
	host := getHost(mc.host)
	port := getPort(mc.port)

	fmt.Printf("%s:%s show\n", host, port)

	if mc.showAll {
		mc.showTrigger = true
		mc.showParameter = true
		mc.showLimit = true
		mc.showEvent = true
		mc.showMeter = true
	}

	client := common.CreateTaklerServiceClient(host, port)
	client.RunQueryShow(
		mc.showTrigger,
		mc.showParameter,
		mc.showLimit,
		mc.showEvent,
		mc.showMeter)

	return nil
}

/*********************************************
	ping
 *********************************************/
type pingCommand struct {
	BaseCommand

	host string
	port string
}

func newPingCommand() *pingCommand {
	c := &pingCommand{}
	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "[query] check the server is running with given host and hort.",
		Long:  "ping server",
		RunE:  c.runCommand,
	}

	pingCmd.Flags().StringVar(&c.host, "host", "", "takler host")
	pingCmd.Flags().StringVar(&c.port, "port", "", "takler port")

	c.cmd = pingCmd
	return c
}

func (mc *pingCommand) runCommand(cmd *cobra.Command, args []string) error {
	host := getHost(mc.host)
	port := getPort(mc.port)

	fmt.Printf("%s:%s ping\n", host, port)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunQueryPing()

	return nil
}
