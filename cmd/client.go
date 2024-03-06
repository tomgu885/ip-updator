package cmd

import (
	"github.com/spf13/cobra"
	"iptables/client"
)

var hearBeatCmd = &cobra.Command{
	Use: "heartbeat",
	Run: func(cmd *cobra.Command, args []string) {
		client.HeartBeat()
	},
}
