package main

import (
	"github.com/Lunarisnia/todo-llm/internal/core"
	lmstudio "github.com/Lunarisnia/todo-llm/internal/core/llm/lm_studio"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

func main() {
	llmModel := lmstudio.NewLMStudio(0.4, -1)
	inputEngine := input.NewInputEngine()

	todoEngine := core.NewTodoEngine(inputEngine, llmModel)
	todoEngine.CreateTask()

	todoEngine.ListTask()
}
