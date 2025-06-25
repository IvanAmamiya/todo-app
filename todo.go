// TEST
package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	panic("unimplemented")
}

func ListTasks() {
	panic("unimplemented")
}

func CompleteTask(id int) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }
    found := false
    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Completed = true
            found = true
            break
        }
    }
    if !found {
        fmt.Printf("Task with ID %d not found.\n", id)
        return
    }
    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
    } else {
        fmt.Printf("Task %d marked as completed.\n", id)
    }
}

func DeleteTask(id int) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }
    newTasks := []Task{}
    found := false
    for _, task := range tasks {
        if task.ID == id {
            found = true
            continue
        }
        newTasks = append(newTasks, task)
    }
    if !found {
        fmt.Printf("Task with ID %d not found.\n", id)
        return
    }
    err = saveTasks(newTasks)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
    } else {
        fmt.Printf("Task %d deleted successfully.\n", id)
    }
}