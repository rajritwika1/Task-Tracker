package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// Task represents a task with details such as ID, description, status, and timestamps
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "todo", "in-progress", or "done"
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const fileName = "tasks.json"

// Helper function to load tasks from JSON file
func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// Helper function to save tasks to JSON file
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

// Function to add a new task
func addTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	return saveTasks(tasks)
}

// Function to update a task's description
func updateTask(id int, description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return saveTasks(tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// Function to delete a task by ID
func deleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveTasks(tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// Function to mark a task's status
func markTask(id int, status string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return saveTasks(tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// Function to list tasks, optionally by status
func listTasks(statusFilter string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	for _, task := range tasks {
		if statusFilter == "" || task.Status == statusFilter {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC1123), task.UpdatedAt.Format(time.RFC1123))
		}
	}
}

// Main function to handle CLI commands
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		description := os.Args[2]
		err := addTask(description)
		if err != nil {
			fmt.Println("Error adding task:", err)
		} else {
			fmt.Println("Task added successfully")
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <description>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		description := os.Args[3]
		err := updateTask(id, description)
		if err != nil {
			fmt.Println("Error updating task:", err)
		} else {
			fmt.Println("Task updated successfully")
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := deleteTask(id)
		if err != nil {
			fmt.Println("Error deleting task:", err)
		} else {
			fmt.Println("Task deleted successfully")
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := markTask(id, "in-progress")
		if err != nil {
			fmt.Println("Error marking task:", err)
		} else {
			fmt.Println("Task marked as in progress")
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := markTask(id, "done")
		if err != nil {
			fmt.Println("Error marking task:", err)
		} else {
			fmt.Println("Task marked as done")
		}
	case "list":
		statusFilter := ""
		if len(os.Args) > 2 {
			statusFilter = os.Args[2]
		}
		listTasks(statusFilter)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add, update, delete, mark-in-progress, mark-done, list")
	}
}
