package storage

import (
	"encoding/json"
	"log"
	"os"
	"tasks/internal/task"
)

// TaskData is a struct for JSON serialization, excluding runtime fields
type TaskData struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Schedule    int    `json:"schedule"`
	Description string `json:"description"`
}

// SaveTasksToJSON writes the current tasks to a JSON file
func SaveTasksToJSON() {
	var taskData []TaskData
	for _, t := range task.Tasks {
		taskData = append(taskData, TaskData{
			ID:          t.ID,
			Name:        t.Name,
			Schedule:    t.Schedule,
			Description: t.Description,
		})
	}
	data, err := json.MarshalIndent(taskData, "", "  ")
	if err != nil {
		log.Println("Error marshaling tasks to JSON:", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		log.Println("Error writing tasks to file:", err)
	}
}

// LoadTasksFromJSON reads tasks from a JSON file and initializes them
func LoadTasksFromJSON() {
	fileContent, err := os.ReadFile("tasks.json")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println("Error reading tasks file:", err)
		}
		return
	}
	var taskData []TaskData
	if err := json.Unmarshal(fileContent, &taskData); err != nil {
		log.Println("Error parsing tasks file:", err)
		return
	}
	for _, data := range taskData {
		t := &task.Task{
			ID:          data.ID,
			Name:        data.Name,
			Schedule:    data.Schedule,
			Description: data.Description,
			TaskFunc:    func() { log.Printf("Executing task '%s': %s", data.Name, data.Description) },
			DoneChan:    make(chan bool),
		}
		task.Tasks[t.ID] = t
		task.StartTaskGoroutine(t)
		if data.ID >= task.NextID {
			task.NextID = data.ID + 1
		}
	}
}
