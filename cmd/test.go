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
		fmt.Println("name:", settings.GetClient().Name)
		fmt.Println("port:", settings.GetClient().LocalPort)
		fmt.Println("type:", settings.GetClient().Type)
		fmt.Println("server:", settings.GetClient().Server)
		return client.Report()
	},
}

var gostRunCmd = &cobra.Command{
	Use: "gost_run",
	Run: func(cmd *cobra.Command, args []string) {
		port := 8001
		if len(args) > 0 {
			var err error
			port, err = strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("failed to convert number:", err.Error())
				return
			}
		}

		nodeIp := "123.0.0.1"
		if len(args) >= 2 {
			nodeIp = args[1]
		}
		updated, err := server.RunGostPortAndIp(port, nodeIp)
		if err != nil {
			fmt.Println("gost_run failed:", err)
			return
		}

		fmt.Printf("updated: %t\n", updated)
		// end of run
	},
}

var gostListCmd = &cobra.Command{
	Use: "gost_list",
	Run: func(cmd *cobra.Command, args []string) {
		port := 8001

		if len(args) > 0 {
			var err error
			port, err = strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("failed to convert number:", err.Error())
				return
			}

		}
		pid, oldIp, err := server.GetPidGostByPort(port)
		if err != nil {
			fmt.Println("failed to getPid:", err)
			return
		}

		fmt.Println("pid:", pid)
		fmt.Println("oldIp:", oldIp)
	},
}
