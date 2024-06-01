package cmd

import (
	"fmt"
	"os/exec"

	"github.com/Akaifox16/gonf/models/spinner"
	"github.com/spf13/cobra"
)

func NewCompleteBranchCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "done [workflow]",
		Short: "Complete source branch and push change to destination branch",
		Args:  cobra.MinimumNArgs(1),
		Run:   executeCompleteBranch,
	}
}

func executeCompleteBranch(cmd *cobra.Command, args []string) {
	workflowName := args[0]

	workflow, found := cfg.Workflows[workflowName]

	if !found {
		fmt.Println("Workflow not found:", workflowName)
		return
	}

	targetBranch := workflow.DefaultBranch
	if targetBranch == "" {
		targetBranch = workflow.DestinationBranch
	}

	runHooks(workflow.PreCommitHooks)
	completeBranch(workflow.SourceBranch, workflow.DestinationBranch, workflow.DestinationRemote)
	runHooks(workflow.PostCommitHooks)
}

func completeBranch(sourceBranch string, destinationBranch string, destinationRemote_optional ...string) {

	fmt.Println("Switching to destination branch and merging changes...")

	// destinationRemote := "origin"
	// if len(destinationRemote_optional) > 0 && destinationRemote_optional[0] != "" {
	// 	destinationRemote = destinationRemote_optional[0]
	// }

	checkoutDestCmd := exec.Command("git", "checkout", destinationBranch)
	if err := checkoutDestCmd.Run(); err != nil {
		fmt.Println("Error checkout to ", destinationBranch, " :", err)
	}

	mergeCmd := exec.Command("git", "merge", sourceBranch)
	if err := spinner.RunSpinner(mergeCmd); err != nil {
		fmt.Println("Error merge ", sourceBranch, " => ", destinationBranch, " : ", err)
	}

	// if err := spinner.RunSpinner("push"); err != nil {
	// 	fmt.Println("Error pushing branch:", err)
	// }
}
