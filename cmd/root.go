/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"iptables/settings"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ip_tools",
	Short: "好玩",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("PersistentPreRunE", cfgFile)
		err := settings.Setup(cfgFile)
		if err != nil {
			return err
		}
		return nil
	},
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.iptables.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "app.yaml", "config file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(hearBeatCmd)
	rootCmd.AddCommand(configTestCmd)
	rootCmd.AddCommand(reportTestCmd)

	rootCmd.AddCommand(natListCmd)
	rootCmd.AddCommand(list2Cmd)
	rootCmd.AddCommand(delRuleCmd)
	rootCmd.AddCommand(addRuleCmd)
	rootCmd.AddCommand(serveCmd)

}
