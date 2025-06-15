package core

import (
	"fmt"
	"log"

	breakdown "github.com/Lunarisnia/todo-llm/internal/core/agents/break_down"
	"github.com/Lunarisnia/todo-llm/internal/core/llm"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

type TodoEngine interface {
	Run()
}

type todoEngineImpl struct {
	State       EngineState
	InputEngine input.InputEngine
	LLMModel    llm.LLM
}

func NewTodoEngine(inputEngine input.InputEngine, llmModel llm.LLM) TodoEngine {
	return &todoEngineImpl{
		State:       MainMenu,
		InputEngine: inputEngine,
		LLMModel:    llmModel,
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
	agent := breakdown.NewBreakDownAgent(t.LLMModel.Clone())
	fmt.Println("1. Add a new todo list")
	fmt.Println("2. List all todo list")
	selection := t.InputEngine.Read("What do you want to do?")
	if selection == "1" {
		prompt := t.InputEngine.Read("What task do you need help?")
		err := agent.Chat(prompt)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (t *todoEngineImpl) changeState(newState EngineState) {
	t.State = newState
}
