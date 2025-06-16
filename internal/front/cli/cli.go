package cli

import (
	"fmt"
	"log"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front"
	mainmenu "github.com/Lunarisnia/todo-llm/internal/front/cli/main_menu"
)

type cliImpl struct {
	State      ProgramState
	TodoEngine core.TodoEngine

	MainMenu mainmenu.MainMenu
}

func NewCLIFrontend(todoEngine core.TodoEngine) front.FrontEnd {
	return &cliImpl{
		State:      MainMenu,
		TodoEngine: todoEngine,
		MainMenu:   *mainmenu.NewMainMenu(todoEngine),
	}
}

func (c *cliImpl) Run() {
	for {
		switch c.State {
		case MainMenu:
			c.MainMenu.Render()
		case ListAll:
			fmt.Println("ListAll")
		case AddNew:
			fmt.Println("AddNew")
		default:
			log.Fatalln("unknown or unimplemented state")
		}
	}
}

func (c *cliImpl) mainMenu() {
	fmt.Println("Main Menu")
}

func (c *cliImpl) SetState(newState ProgramState) {
	c.State = newState
}
