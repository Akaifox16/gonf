package cmd

import (
	"fmt"

	"github.com/Akaifox16/gonf/config"
	"github.com/Akaifox16/gonf/spinner"
	"github.com/spf13/cobra"
)

func NewSyncBranchCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "sync [workflow]",
		Short: "Sync change from remote repository",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			workflowName := args[0]

			workflow, found := cfg.Workflows[workflowName]

			if !found {
				fmt.Println("Workflow not found:", workflowName)
				return
			}

			syncBranch(workflow)

		},
	}
}

func syncBranch(workflow config.Workflow) {
	fmt.Println("Syncing code from the default branch...")
	if err := spinner.RunSpinner("pull"); err != nil {
		fmt.Println("Error syncing branch:", err)
	}
}
