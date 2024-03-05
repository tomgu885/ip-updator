package cmd

import (
	"github.com/spf13/cobra"
	"iptables/client"
)

var hearBeatCmd = &cobra.Command{
	Use: "client_heartbeat",
	Run: func(cmd *cobra.Command, args []string) {
		client.HeartBeat()
	},
}
