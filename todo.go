// test comment
package main

import (
	"encoding/json"
	"fmt"
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
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		if err.Error() == "EOF" {
			return []Task{}, nil
		}
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
	tasks, err := loadTasks()
	if err != nil {
		println("Error loading tasks:", err.Error())
		return
	}
	task := Task{
		ID:    nextID(tasks),
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	if err := saveTasks(tasks); err != nil {
		println("Error saving tasks:", err.Error())
		return
	}
	println("Added task:", title)
}

func ListTasks() {
	tasks, err := loadTasks()
	if err != nil {
		println("Error loading tasks:", err.Error())
		return
	}
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		println(t.ID, ":", t.Title, status)
	}
}

func CompleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		println("Error loading tasks:", err.Error())
		return
	}
	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}
	if !found {
		println("Task not found:", id)
		return
	}
	if err := saveTasks(tasks); err != nil {
		println("Error saving tasks:", err.Error())
		return
	}
	println("Task completed:", id)
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
