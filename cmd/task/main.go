package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mohammad-alj/task/iotask"
)

func main() {
	commandsMap := map[string]string{
		"help":                 "Show this message",
		"add":                  "Create a new task",
		"list":                 "List all tasks",
		"remove [id] [id] ...": "Remove specefic tasks for completion",
	}

	args := os.Args[1:]
	argsLength := len(args)

	help := func() {
		fmt.Println("usage: task [COMMAND]")
		fmt.Println()

		for cmd, desc := range commandsMap {
			fmt.Printf("    %s: %s\n", cmd, desc)
		}
	}

	if 0 == argsLength {
		help()
		return
	}

	command := args[0]

	switch command {
	case "help":
		help()
	case "add":
		fmt.Print("New Task: ")

		reader := bufio.NewReader(os.Stdin)
		taskContent, err := reader.ReadString('\n')
		taskContent = strings.TrimRight(taskContent, "\n")
		if err != nil {
			panic(err)
		}

		_, err = iotask.WriteTask(taskContent)
		if err != nil {
			panic(err)
		}

		fmt.Println("Task added")

	case "list":
		tasks, err := iotask.GetTasks()
		if err != nil {
			panic(err)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks!")
			return
		}

		for _, task := range tasks {
			fmt.Printf("ID: %d, Content: %s\n", task.Id, task.Content)
		}
	case "remove":
		for _, arg := range args[1:] {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Couldn't delete task with id", id)
				return
			}

			task, err := iotask.RemoveTask(id)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Printf("Task \"%s\" removed\n", task.Content)
		}
		tasks, err := iotask.GetTasks()
		if err != nil {
			panic(err)
		}

		iotask.RerangeIDs(tasks)
		iotask.WriteTasks(tasks)
	default:
		help()
	}
}
