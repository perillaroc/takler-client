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
}

func newShowCommand() *showCommand {
	c := &showCommand{}
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show",
		Long:  "show",
		RunE:  c.runCommand,
	}

	showCmd.Flags().StringVar(&c.host, "host", "", "takler host")
	showCmd.Flags().StringVar(&c.port, "port", "", "takler port")

	c.cmd = showCmd
	return c
}

func (mc *showCommand) runCommand(cmd *cobra.Command, args []string) error {
	host := getHost(mc.host)
	port := getPort(mc.port)

	fmt.Printf("%s:%s show\n", host, port)

	client := common.CreateTaklerServiceClient(host, port)
	client.RunQueryShow()

	return nil
}
