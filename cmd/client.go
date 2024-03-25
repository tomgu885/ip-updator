package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iptables/client"
	"iptables/settings"
)

var hearBeatCmd = &cobra.Command{
	Use: "heartbeat",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start client heartbeat")
		fmt.Println("name:", settings.GetClient().Name)
		fmt.Println("port:", settings.GetClient().LocalPort)
		client.HeartBeat()
	},
}
