package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Akaifox16/gonf/cmd"
	"github.com/Akaifox16/gonf/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "gonf",
	Short: "Git Configurable Extension",
	Long:  `A configurable Git extension CLI for your desire workflow.`,
}

func init() {
	homeDir, exists := os.LookupEnv("HOME")
	if !exists {
		fmt.Println("HOME environment variable is not set")
		return
	}
	configFilePath := filepath.Join(homeDir, ".config", "gonf", "config.toml")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", configFilePath, "config file (default is $HOME/.config/gonf/config.toml)")
}

func main() {
	var err error
	cfg, err = config.LoadConfig(cfgFile)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	rootCmd.AddCommand(cmd.NewBranchCommand(cfg))
	rootCmd.AddCommand(cmd.NewSyncBranchCommand(cfg))

	rootCmd.Execute()
}
