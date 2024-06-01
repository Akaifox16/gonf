package cmd

import (
	"fmt"
	"os/exec"
)

func runHooks(hook string) {
	if hook == "" {
		return
	}

	cmd := exec.Command("sh", "-c", hook)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error runing hook:", err)
	}
}
