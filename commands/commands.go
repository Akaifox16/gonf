package commands

import (
	"fmt"
	"os/exec"

	"github.com/Akaifox16/gonf/branchinput"
	"github.com/Akaifox16/gonf/config"
	"github.com/Akaifox16/gonf/spinner"
)

var cfg *config.Config

func SetConfig(config *config.Config) {
	cfg = config
}

func CreateBranch(workflow config.Workflow, branchName string) {
	if branchName == "" {
		branchName = branchinput.OpenBranchTextInput()
	}

	cmd := exec.Command("git", "checkout", "-b", branchName)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error creating branch:", err)
	}
}

func SyncBranch(workflow config.Workflow) {
	fmt.Println("Syncing code from the default branch...")
	if err := spinner.RunSpinner("pull"); err != nil {
		fmt.Println("Error syncing branch:", err)
	}
}

func CompleteBranch(workflow config.Workflow, branchName string) {
	runHooks(workflow.PreCommitHooks)
	fmt.Println("Switching to destination branch and merging changes...")
	if err := spinner.RunSpinner("push"); err != nil {
		fmt.Println("Error pushing branch:", err)
	}
	runHooks(workflow.PostCommitHooks)
}

func runHooks(hook string) {
	if hook == "" {
		return
	}

	cmd := exec.Command("sh", "-c", hook)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error runing hook:", err)
	}
}
