package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aquemaati/lamoul01/internal/watcher"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("\033[31mUsage: lamoul01 <language>\033[0m")
		return
	}

	language := args[1]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("\033[31mError retrieving home directory: %v\033[0m\n", err)
		return
	}

	switch language {
	case "js":
		testDir := filepath.Join(homeDir, ".lamoul01", "public", "js", "tests", "test.mjs")
		fmt.Printf("🤡\033[32mRunning JavaScript tests in %s...\033[0m\n", testDir)
		watcher.Watcher(runJSTests, testDir)
	case "go":
		testDir := filepath.Join(homeDir, ".lamoul01", "public", "go", "tests")
		fmt.Printf("🤡\033[32mRunning Go tests in %s...\033[0m\n", testDir)
		watcher.Watcher(runGoTests, testDir)
	default:
		fmt.Println("🤢\033[31mUnknown language:\033[0m", language)
	}
}

func runJSTests(exName, folder, testPath string) {
	fmt.Printf("🤪\033[34mRunning JavaScript tests for: %s\033[0m\n", exName)
	cmd := exec.Command("node", testPath, folder, exName)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("🤢\033[31mError running JavaScript tests: %v\033[0m\n", err)
	}
}

func runGoTests(exName, folder, testPath string) {
	fmt.Printf("🤪\033[34mRunning Go tests for: %s\033[0m\n", exName)
}
