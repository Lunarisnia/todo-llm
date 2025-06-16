package main

import (
	"fmt"

	"github.com/Lunarisnia/todo-llm/internal/core"
	lmstudio "github.com/Lunarisnia/todo-llm/internal/core/llm/lm_studio"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

func main() {
	llmModel := lmstudio.NewLMStudio(0.4, -1)
	inputEngine := input.NewInputEngine()

	todoEngine := core.NewTodoEngine(inputEngine, llmModel)
	newTask, err := todoEngine.CreateTask("Make an omellete")
	if err != nil {
		panic(err)
	}
	tasks, err := todoEngine.HelpBreakdown(newTask)
	if err != nil {
		panic(err)
	}
	for _, t := range tasks {
		newTask.AddSubtask(&t)
	}

	storedTasks, err := todoEngine.ListTask()
	for _, t := range storedTasks {
		fmt.Println("-", t.Content)
		for _, ct := range t.Subtask {
			fmt.Println("\t-", ct.Content)
		}
	}

}
