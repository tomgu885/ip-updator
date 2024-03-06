package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iptables/logger"
	"iptables/model"
	"iptables/server"
	"strconv"
)

var natListCmd = &cobra.Command{
	Use:   "nat_list",
	Short: "nat rules list",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("nat rule list")
		ruleType := "nat"
		if len(args) > 0 {
			ruleType = args[0]
		}

		rules, err := server.ListRules(ruleType)
		if err != nil {
			logger.Errorf("failed to get rules:%v", err)
			return
		}

		fmt.Printf("rules: %v\n", rules)
	},
}

var list2Cmd = &cobra.Command{
	Use: "list2",
	Run: func(cmd *cobra.Command, args []string) {
		rules, err := server.ListNatWithNumber()
		if err != nil {
			logger.Errorf("list2 failed: %v", err)
			return
		}

		for _, rule := range rules {
			fmt.Printf("rule:%d listen:%d, to:%s port:%d\n", rule.RuleNum, rule.ListenPort, rule.ToIp, rule.ToPort)
		}

		fmt.Printf("finished\n")
	},
}

var addRuleCmd = &cobra.Command{
	Use:   "rule_add",
	Short: "rule_add ip port",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("no rule")
			return
		}

		addr := args[0]
		port, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("port should be number")
			return
		}
		rule := model.NatRule{
			RuleNum:    0,
			ListenPort: port,
			ToIp:       addr,
			ToPort:     port,
		}
		deleted, err := server.DelNatPortByToPort(rule.ToPort)

		if err != nil {
			logger.Errorf("DelNatPortByToPort failed: %v", err)
			return
		}

		if deleted {
			logger.Infof("rule updated")
		}

		err = server.AddNatRule(rule)

		if err != nil {
			logger.Errorf("failed to add rule:%v", err)
			return
		}

		return
	},
}

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}
