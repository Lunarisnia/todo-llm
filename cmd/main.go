package main

import (
	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

func main() {
	inputEngine := input.NewInputEngine()

	todoEngine := core.NewTodoEngine(inputEngine)
	todoEngine.Run()
}
