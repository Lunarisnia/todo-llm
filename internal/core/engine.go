package core

import (
	"fmt"
	"log"

	breakdown "github.com/Lunarisnia/todo-llm/internal/core/agents/break_down"
	"github.com/Lunarisnia/todo-llm/internal/core/agents/describe"
	"github.com/Lunarisnia/todo-llm/internal/core/llm"
	"github.com/Lunarisnia/todo-llm/internal/core/task"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

type TodoEngine interface {
	CreateTask(content string) error
	CreateSubtask(taskId uint, content string) error
	ListTask() error
}

// TODO: Connect it to a database
type todoEngineImpl struct {
	State       EngineState
	InputEngine input.InputEngine
	Storage     []task.Task

	LLMModel       llm.LLM
	TaskAgent      *breakdown.BreakDownAgent
	DescriberAgent *describe.DescribeAgent
}

func NewTodoEngine(inputEngine input.InputEngine, llmModel llm.LLM) TodoEngine {
	taskAgent := breakdown.NewBreakDownAgent(llmModel.Clone())
	describerAgent := describe.NewDescribeAgent(llmModel.Clone())
	todoEngine := &todoEngineImpl{
		State:          MainMenu,
		InputEngine:    inputEngine,
		LLMModel:       llmModel,
		TaskAgent:      taskAgent,
		DescriberAgent: describerAgent,
	}
	return todoEngine
}

func (t *todoEngineImpl) CreateTask(content string) error {
	// TODO: Create task
	return nil
}

func (t *todoEngineImpl) CreateSubtask(taskId uint, content string) error {
	// TODO: Create subtask
	return nil
}

func (t *todoEngineImpl) ListTask() error {
	// TODO: Create list task
	return nil
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

// TODO: Make this a proper main menu
func (t *todoEngineImpl) mainMenu() {
	fmt.Println("1. Add a new todo list")
	fmt.Println("2. List all todo list")
	selection := t.InputEngine.Read("What do you want to do?")
	described, err := t.DescriberAgent.Chat(selection)
	if err != nil {
		return
	}
	fmt.Println(described.Choices)
	result, err := t.TaskAgent.Chat(described.Choices[0].Message.Content)
	if err != nil {
		return
	}
	fmt.Println(result.Choices)

}

func (t *todoEngineImpl) changeState(newState EngineState) {
	t.State = newState
}
