package cmd

import (
	"fmt"
	"os/exec"

	"github.com/Akaifox16/gonf/branchinput"
	"github.com/Akaifox16/gonf/config"
	"github.com/spf13/cobra"
)

func NewBranchCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "new [workflow] [branch name]",
		Short: "Create a new branch",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			workflowName := args[0]
			branchName := ""
			if len(args) > 1 {
				branchName = args[1]
			}

			_, found := cfg.Workflows[workflowName]

			if !found {
				fmt.Println("Workflow not found:", workflowName)
				return
			}

			createBranch(branchName)
		},
	}
}

func createBranch(branchName string) {
	if branchName == "" {
		branchName = branchinput.OpenBranchTextInput()
	}

	cmd := exec.Command("git", "checkout", "-b", branchName)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error creating branch:", err)
	}
}
