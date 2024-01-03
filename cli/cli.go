package cli

import (
	"fmt"

	"github.com/mohammad-alj/task/taskdata"
)

func PrintHelpMenu() {
	commandsMap := map[string]string{
		"help":            "Show this message",
		"add":             "Create a new task",
		"list":            "List all tasks",
		"remove [id] ...": "Remove specefic tasks for completion",
	}

	longestMapKey := 0
	for k := range commandsMap {
		keyLen := len(k)
		if keyLen > longestMapKey {
			longestMapKey = keyLen
		}
	}

	fmt.Println("usage: task [COMMAND]")
	fmt.Println()

	for cmd, desc := range commandsMap {
		spaces := longestMapKey - len(cmd) + 2
		fmt.Printf("    %s", cmd)
		for i := 0; i < spaces; i++ {
			fmt.Print(" ")
		}
		fmt.Println(desc)
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
