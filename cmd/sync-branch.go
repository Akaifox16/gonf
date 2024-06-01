package cmd

import (
	"fmt"
	"os/exec"

	"github.com/Akaifox16/gonf/models/spinner"
	"github.com/spf13/cobra"
)

func NewSyncBranchCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sync [workflow]",
		Short: "Sync change from remote repository",
		Args:  cobra.MinimumNArgs(1),
		Run:   executeSyncBranch,
	}
}

func executeSyncBranch(cmd *cobra.Command, args []string) {
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

	remote := workflow.DefaultBranchRemote
	if remote == "" {
		remote = workflow.DestinationRemote
	}

	syncBranch(targetBranch, remote)

}

func syncBranch(targetBranch string, remoteName_optional ...string) {

	fmt.Println("Syncing code from the default branch...")

	remoteName := "origin"

	if len(remoteName_optional) > 0 && remoteName_optional[0] != "" {
		remoteName = remoteName_optional[0]
	}
	// fmt.Printf("remote: %v, branch: %v\n", remoteName, targetBranch)

	pullCmd := exec.Command("git", "pull", remoteName, targetBranch)
	// fmt.Printf("Prepare Cmd: %+v\n", pullCmd)

	if err := spinner.RunSpinner(pullCmd); err != nil {
		fmt.Println("Error syncing branch:", err)
	}
}
