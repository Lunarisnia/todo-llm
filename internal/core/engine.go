package core

import (
	"fmt"

	"github.com/Lunarisnia/todo-llm/internal/input"
)

type TodoEngine interface {
	Run()
}

type todoEngineImpl struct {
	InputEngine input.InputEngine
}

func NewTodoEngine(inputEngine input.InputEngine) TodoEngine {
	return &todoEngineImpl{
		InputEngine: inputEngine,
	}
}

func (t *todoEngineImpl) Run() {
	fmt.Println("1. Add a new todo list")
	fmt.Println("2. List all todo list")
	t.InputEngine.Read("What do you want to do?")
}
