package cli

import (
	"fmt"

	"github.com/mohammad-alj/task/taskdata"
)

type command struct {
	Name        string
	Description string
}

func PrintHelpMenu() {
	commands := []command{
		{
			Name:        "add",
			Description: "Create a new task",
		},
		{
			Name:        "list",
			Description: "List all tasks",
		},
		{
			Name:        "remove <id>",
			Description: "Remove tasks by inserting their ID's",
		},
	}

	longestCmd := 0
	for _, cmd := range commands {
		cmdLen := len(cmd.Name)
		if cmdLen > longestCmd {
			longestCmd = cmdLen
		}
	}

	fmt.Println("usage: task [COMMAND]")
	fmt.Println()

	for _, cmd := range commands {
		spaces := longestCmd - len(cmd.Name) + 2
		fmt.Printf("    %s", cmd.Name)
		for i := 0; i < spaces; i++ {
			fmt.Print(" ")
		}
		fmt.Println(cmd.Description)
	}
}

func ListTasks(tasks []taskdata.Task) {
	if len(tasks) == 0 {
		fmt.Println("You have no tasks!")
		return
	}
	for _, task := range tasks {
		fmt.Printf("Task %d: %s\n", task.Id, task.Content)
	}
}
