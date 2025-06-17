package mainmenu

import (
	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/clistate"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
)

type MainMenu struct {
	CursorIndex int
	Options     []menu.Option
}

func NewMainMenu(todoEngine core.TodoEngine) menu.Menu {
	mainMenu := &MainMenu{
		CursorIndex: 0,
		Options:     make([]menu.Option, 0),
	}

	mainMenu.Options = append(mainMenu.Options, menu.NewOption(
		"Add New Task",
		mainMenu.AddNewTask,
	))

	mainMenu.Options = append(mainMenu.Options, menu.NewOption(
		"List All Task",
		mainMenu.ListAllTask,
	))

	return mainMenu
}

func (m *MainMenu) GetCursorIndex() int {
	return m.CursorIndex
}

func (m *MainMenu) SetCursorIndex(newIndex int) {
	m.CursorIndex = newIndex
}

func (m *MainMenu) GetOptions() []menu.Option {
	return m.Options
}

func (m *MainMenu) Callback(index int) error {
	err := m.Options[index].Action()
	if err != nil {
		return err
	}
	return nil
}

func (m *MainMenu) AddNewTask() error {
	clistate.State = clistate.AddNew
	return nil
}

func (m *MainMenu) ListAllTask() error {
	clistate.State = clistate.ListAll
	return nil
}
