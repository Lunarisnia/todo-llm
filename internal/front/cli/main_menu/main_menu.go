package mainmenu

import (
	"fmt"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/input"
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

// NEXT TODO: Rendering engine
// TODO: Should make a rendering engine to handle all the rendering
func (m *MainMenu) Render() {
	fmt.Print("\033[H\033[2J")
	for i, opt := range m.Options {
		prefix := ""
		if m.CursorIndex == i {
			prefix = "*"
		}
		fmt.Println(prefix, opt.Title)
	}
	r, _ := input.ReadKeyboard("")
	if r == 'j' {
		m.CursorIndex++
	}
	if r == 'k' {
		m.CursorIndex--
		if m.CursorIndex < 0 {
			m.CursorIndex = len(m.Options) - 1
		}
	}
	m.CursorIndex = m.CursorIndex % len(m.Options)
}

func (m *MainMenu) AddNewTask() error {
	return nil
}

func (m *MainMenu) ListAllTask() error {
	return nil
}
