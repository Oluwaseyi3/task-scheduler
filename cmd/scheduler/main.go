package main

import (
	"fmt"
	"os"
	"tasks/internal/cli"
	"tasks/internal/storage"
)

func main() {
	// Load existing tasks from JSON file
	storage.LoadTasksFromJSON()

	// Start the interactive CLI
	fmt.Println("Task Scheduler started. Type 'help' for commands.")
	cli.StartCLI()

	// Handle shutdown
	fmt.Println("Shutting down...")
	cli.Shutdown()
	os.Exit(0)
}
