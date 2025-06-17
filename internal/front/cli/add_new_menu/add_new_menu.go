package addnewmenu

import (
	"fmt"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/clistate"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

type AddNewMenu struct {
	CursorIndex int
	Options     []menu.Option
	TodoEngine  core.TodoEngine
}

func NewAddNewMenu(todoEngine core.TodoEngine) menu.Menu {
	addNewMenu := &AddNewMenu{
		CursorIndex: 0,
		Options:     make([]menu.Option, 0),
		TodoEngine:  todoEngine,
	}

	addNewMenu.Options = append(addNewMenu.Options, menu.NewOption(
		"Enter the name of your new task:",
		addNewMenu.AddNewTask,
	))

	return addNewMenu
}

func (a *AddNewMenu) GetCursorIndex() int {
	return a.CursorIndex
}

func (a *AddNewMenu) SetCursorIndex(newIndex int) {
	a.CursorIndex = newIndex
}

func (a *AddNewMenu) GetOptions() []menu.Option {
	return a.Options
}

func (a *AddNewMenu) Callback(index int) error {
	return a.Options[index].Action()
}

func (a *AddNewMenu) AddNewTask() error {
	userInput := input.Read("")
	task, err := a.TodoEngine.CreateTask(userInput)
	if err != nil {
		return err
	}
	wantToBreakDown := input.Read("Do you want me to help break this task down for you? (y/n):")
	if wantToBreakDown == "y" {
		subtasks, err := a.TodoEngine.HelpBreakdown(task)
		if err != nil {
			return err
		}
		for _, subtask := range subtasks {
			fmt.Println("\t-", subtask.Content)
		}
		addThisOrNot := input.Read("Add these as a subtask? (y/n):")
		if addThisOrNot == "y" {
			for _, subtask := range subtasks {
				task.AddSubtask(&subtask)
			}
			fmt.Println("Added as subtasks")
			input.Read("")
		} else {
			fmt.Println("Will not add this as a subtask")
			input.Read("")
		}
	} else {
		fmt.Println("okay")
		input.Read("")
	}
	clistate.State = clistate.MainMenu
	return nil
}
