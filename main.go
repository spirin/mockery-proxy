package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/vektra/mockery/v2/cmd"
)

var passthroughCommands = map[string]bool{
	"completion": true,
	"help":       true,
	"init":       true,
	"migrate":    true,
	"showconfig": true,
	"version":    true,
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		runMockery(args)
		return
	}

	if strings.HasPrefix(args[0], "--config") || passthroughCommands[args[0]] {
		runMockery(args)
		return
	}

	cmd.Execute()
}

func runMockery(args []string) {
	cmd := exec.Command("mockery", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "mockery failed: %v\n", err)
		os.Exit(1)
	}
}
