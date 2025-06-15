package core

import (
	"fmt"
	"log"

	"github.com/Lunarisnia/todo-llm/internal/input"
)

type TodoEngine interface {
	Run()
}

type todoEngineImpl struct {
	State       EngineState
	InputEngine input.InputEngine
}

func NewTodoEngine(inputEngine input.InputEngine) TodoEngine {
	return &todoEngineImpl{
		State:       MainMenu,
		InputEngine: inputEngine,
	}
}

func (t *todoEngineImpl) Run() {
	for {
		switch t.State {
		case MainMenu:
			t.mainMenu()
		case AddNew:
			log.Fatalln("TODO: Implement add new")
		case ListAll:
			log.Fatalln("TODO: Implement list all")
		default:
			log.Fatalln("unknown or unimplemented state")
		}
	}
}

func (t *todoEngineImpl) mainMenu() {
	fmt.Println("1. Add a new todo list")
	fmt.Println("2. List all todo list")
	t.InputEngine.Read("What do you want to do?")
}

func (t *todoEngineImpl) changeState(newState EngineState) {
	t.State = newState
}
