package core

import (
	breakdown "github.com/Lunarisnia/todo-llm/internal/core/agents/break_down"
	"github.com/Lunarisnia/todo-llm/internal/core/agents/describe"
	"github.com/Lunarisnia/todo-llm/internal/core/llm"
	lmstudio "github.com/Lunarisnia/todo-llm/internal/core/llm/lm_studio"
	"github.com/Lunarisnia/todo-llm/internal/core/task"
)

type TodoEngine interface {
	CreateTask(content string) (*task.Task, error)
	HelpBreakdown(tsk *task.Task) ([]task.Task, error)
	ListTask() ([]*task.Task, error)
}

// TODO: Connect it to a database
type todoEngineImpl struct {
	Storage []*task.Task

	LLMModel       llm.LLM
	TaskAgent      *breakdown.BreakDownAgent
	DescriberAgent *describe.DescribeAgent
}

func NewTodoEngine(llmModel llm.LLM) TodoEngine {
	taskAgent := breakdown.NewBreakDownAgent(llmModel.Clone())
	describerAgent := describe.NewDescribeAgent(llmModel.Clone())
	todoEngine := &todoEngineImpl{
		LLMModel:       llmModel,
		TaskAgent:      taskAgent,
		DescriberAgent: describerAgent,
	}
	return todoEngine
}

func (t *todoEngineImpl) CreateTask(content string) (*task.Task, error) {
	newTask := task.NewTask(content, nil)
	t.Storage = append(t.Storage, &newTask)
	return &newTask, nil
}

func (t *todoEngineImpl) HelpBreakdown(tsk *task.Task) ([]task.Task, error) {
	describedProblem, err := t.DescriberAgent.Chat(tsk.Content)
	if err != nil {
		return nil, err
	}
	problemBrokeDown, err := t.TaskAgent.Chat(describedProblem.Choices[0].Message.Content)
	if err != nil {
		return nil, err
	}
	parsedTasks := lmstudio.ParseBulletPoint(problemBrokeDown)

	tasks := make([]task.Task, 0)
	for _, pt := range parsedTasks {
		tasks = append(tasks, task.NewTask(pt, nil))
	}

	return tasks, nil
}

func (t *todoEngineImpl) ListTask() ([]*task.Task, error) {
	return t.Storage, nil
}

// TODO:  Deleting todo list
