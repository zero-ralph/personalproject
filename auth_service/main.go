package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zero-ralph/personalporject_users/auth_service/pkg/config"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	fmt.Println("=========================================")
	fmt.Println("=========================================")
	fmt.Println("==         System is Starting          ==")
	fmt.Println("=========================================")
	fmt.Println("=========================================")

	flag.Usage = usage
	configFile := flag.String("settings", "config.toml", "Set your config.toml")
	flag.Parse()

	configManager := config.NewConfigManager()

	if err := configManager.ReadConfigFile(*configFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server := &config.Server{}
	server.Config = configManager

	server.ExecuteStart()

}
