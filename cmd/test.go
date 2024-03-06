package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iptables/client"
	"iptables/server"
	"iptables/settings"
	"strconv"
)

var delRuleCmd = &cobra.Command{
	Use: "rule_del",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("rule number please")
		}

		ruleNum, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("faile: err:", err.Error())
		}

		err = server.DelNatByLineNumber(ruleNum)

		if err != nil {
			fmt.Println("faile2: err:", err.Error())
			return
		}

		fmt.Println("finished")
	},
}

var configTestCmd = &cobra.Command{
	Use:   "config_test",
	Short: "config test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("secret:", settings.GetGlobal().Secret)
	},
}

var reportTestCmd = &cobra.Command{
	Use:   "report",
	Short: "report test",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("reporting...")
		return client.Report()
	},
}
