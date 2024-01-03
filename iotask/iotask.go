package iotask

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"
)

type Task struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func RerangeIDs(tasks []Task) {
	for i := 0; i < len(tasks); i++ {
		tasks[i].Id = i + 1
	}
}

func getTasksPath() string {
	user, _ := user.Current()
	filePath := path.Join(user.HomeDir, ".tasks.json")

	// Create the file if it doesn't exist
	if bytes, err := os.ReadFile(filePath); len(bytes) == 0 || err != nil {
		os.WriteFile(filePath, []byte("[]"), 0666)
	}

	return filePath
}

func GetTasks() ([]Task, error) {
	tasksBytes, err := os.ReadFile(getTasksPath())
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(tasksBytes, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteTask(taskContent string) (*Task, error) {
	if taskContent == "" {
		return nil, errors.New("Task content is empty")
	}

	tasks, err := GetTasks()
	if err != nil {
		return nil, err
	}

	newTask := Task{
		Id:      len(tasks) + 1,
		Content: taskContent,
	}

	tasks = append(tasks, newTask)
	tasksBytes, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(getTasksPath(), tasksBytes, 0666)
	if err != nil {
		return nil, err
	}

	return &newTask, err
}

func WriteTasks(tasks []Task) error {
	tasksBytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(getTasksPath(), tasksBytes, 0666)
	if err != nil {
		return err
	}

	return nil
}

func RemoveTask(taskId int) (*Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return nil, err
	}

	var deletedTask *Task = nil
	for i, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			deletedTask = &task
			break
		}
	}
	if deletedTask == nil {
		return nil, errors.New(fmt.Sprintf("Task with the ID %d was not found", taskId))
	}

	tasksBytes, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}

	os.WriteFile(getTasksPath(), tasksBytes, 0666)

	return deletedTask, nil
}
