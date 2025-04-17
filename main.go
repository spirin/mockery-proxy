package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	selfName := filepath.Base(os.Args[0])

	if len(args) == 0 {
		runMockery(selfName, args)
		return
	}

	if strings.HasPrefix(args[0], "--config") || passthroughCommands[args[0]] {
		runMockery(selfName, args)
		return
	}

	// Default to v2 logic
	cmd.Execute()
}

func runMockery(selfName string, args []string) {
	var target string
	switch selfName {
	case "mockery":
		target = "mockery3"
	case "mockery-proxy":
		target = "mockery"
	default:
		// fallback if renamed differently
		target = "mockery3"
	}

	cmd := exec.Command(target, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %v\n", target, err)
		os.Exit(1)
	}
}
