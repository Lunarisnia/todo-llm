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
	TaskAgent   *breakdown.BreakDownAgent
}

func NewTodoEngine(inputEngine input.InputEngine, llmModel llm.LLM) TodoEngine {
	taskAgent := breakdown.NewBreakDownAgent(llmModel.Clone())
	todoEngine := &todoEngineImpl{
		State:       MainMenu,
		InputEngine: inputEngine,
		LLMModel:    llmModel,
		TaskAgent:   taskAgent,
	}
	return todoEngine
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
	selection := t.InputEngine.Read("What do you want to do?")
	result, err := t.TaskAgent.Chat(selection)
	if err != nil {
		return
	}
	fmt.Println(result.Choices)
}

func (t *todoEngineImpl) changeState(newState EngineState) {
	t.State = newState
}
