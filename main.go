/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"iptables/cmd"
	"iptables/logger"
)

func main() {
	//configPath := "app.ini"

	//settings.Setup(configPath)

	defer logger.Sync()
	cmd.Execute()
}
