package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"tasks/internal/storage"
	"tasks/internal/task"
)

var nextID int32 = 1

func getNextID() int {
	return int(atomic.AddInt32(&nextID, 1)) - 1
}

func AddTask(name string, schedule int, description string) {
	t := &task.Task{
		ID:          getNextID(),
		Name:        name,
		Schedule:    schedule,
		Description: description,
		TaskFunc: func() {
			fmt.Printf("Executing task '%s' : %s\n", name, description)
		},
		DoneChan: make(chan bool),
	}

	task.Tasks[t.ID] = t
	task.StartTaskGoroutine(t)
	storage.SaveTasksToJSON()
	fmt.Printf("✅ Added task %d: %s\n", t.ID, t.Name)
}

func DeleteTask(id int) {
	if t, exists := task.Tasks[id]; exists {
		select {
		case <-t.DoneChan:
		default:
			close(t.DoneChan)
		}
		delete(task.Tasks, id)
		storage.SaveTasksToJSON()
		fmt.Printf("✅ Deleted task %d\n", id)
	} else {
		fmt.Println("❌ Task not found")
	}
}

func ListTasks() {
	if len(task.Tasks) == 0 {
		fmt.Println("ℹ️ No tasks scheduled")
		return
	}
	fmt.Println("📋 Scheduled Tasks:")
	for _, t := range task.Tasks {
		fmt.Printf("🔹 ID %d | Name: %s | Schedule: %d min | Desc: %s\n", t.ID, t.Name, t.Schedule, t.Description)
	}
}

func processCommand(input string) {

	// Ensure we remove \r and extra spaces
	input = strings.TrimSpace(strings.ReplaceAll(input, "\r", ""))

	// **Fix: Properly extract quoted arguments**
	re := regexp.MustCompile(`"([^"]*)"|(\S+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	// Convert match groups into a list of extracted arguments
	var args []string
	for _, match := range matches {
		if match[1] != "" {
			args = append(args, match[1]) // Quoted argument
		} else {
			args = append(args, match[2]) // Regular word
		}
	}

	if len(args) == 0 {
		fmt.Println("❌ No input detected")
		return
	}

	command := strings.ToLower(strings.TrimSpace(args[0]))

	switch command {
	case "add":
		fmt.Println("✅ Command recognized: 'add'") // Debugging
		if len(args) < 4 {
			fmt.Println("❌ Usage: add <name> <schedule> <description>")
			return
		}
		name := args[1]
		schedule, err := strconv.Atoi(args[2])
		if err != nil || schedule <= 0 {
			fmt.Println("❌ Invalid schedule: must be a positive integer")
			return
		}
		description := args[3]

		AddTask(name, schedule, description)

	case "delete":
		fmt.Println("✅ Command recognized: 'delete'") // Debugging
		if len(args) < 2 {
			fmt.Println("❌ Usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("❌ Invalid ID")
			return
		}
		DeleteTask(id)

	case "list":
		fmt.Println("✅ Command recognized: 'list'") // Debugging
		ListTasks()

	case "help":
		fmt.Println("✅ Command recognized: 'help'") // Debugging
		fmt.Println("📌 Commands:")
		fmt.Println("➕ add <name> <schedule> <desc>  - Add a new task (schedule in minutes)")
		fmt.Println("❌ delete <id>  - Delete a task by ID")
		fmt.Println("📋 list  - List all tasks")
		fmt.Println("ℹ️ help  - Show this help")
		fmt.Println("🚪 exit  - Exit the scheduler")

	default:
		fmt.Printf("❌ Unknown command received: '%s'\n", command)
	}
}

func StartCLI() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println("❌ Error reading input")
			return
		}
		input := strings.TrimSpace(scanner.Text()) // Read and clean input
		if input == "exit" {
			return
		}
		processCommand(input)
	}
}

func Shutdown() {
	for _, t := range task.Tasks {
		select {
		case <-t.DoneChan:
		default:
			close(t.DoneChan)
		}
	}
}
