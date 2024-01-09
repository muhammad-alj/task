package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mohammad-alj/task/cli"
	"github.com/mohammad-alj/task/taskdata"
)

func main() {
	args := os.Args[1:]
	argsLength := len(args)

	if 0 == argsLength {
		cli.PrintHelpMenu()
		return
	}

	command := args[0]

	switch command {
	case "add":
		fmt.Print("New Task: ")

		reader := bufio.NewReader(os.Stdin)
		taskContent, err := reader.ReadString('\n')
		taskContent = strings.TrimRight(taskContent, "\n")
		if err != nil {
			panic(err)
		}

		_, err = taskdata.WriteTask(taskContent)
		if err != nil {
			panic(err)
		}

		fmt.Println("Task added")

	case "list":
		tasks, err := taskdata.GetTasks()
		if err != nil {
			panic(err)
		}
		cli.ListTasks(tasks)
	case "remove":
		for _, arg := range args[1:] {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Couldn't delete task with id", id)
				return
			}

			task, err := taskdata.RemoveTask(id)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Printf("Task \"%s\" removed\n", task.Content)
		}
		tasks, err := taskdata.GetTasks()
		if err != nil {
			panic(err)
		}

		taskdata.RerangeIDs(tasks)
		taskdata.WriteTasks(tasks)
	default:
		cli.PrintHelpMenu()
	}
}
