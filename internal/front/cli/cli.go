package cli

import (
	"fmt"
	"log"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front"
)

type cliImpl struct {
	State      ProgramState
	TodoEngine core.TodoEngine
}

func NewCLIFrontend(todoEngine core.TodoEngine) front.FrontEnd {
	return &cliImpl{
		State:      MainMenu,
		TodoEngine: todoEngine,
	}
}

func (c *cliImpl) Run() {
	for {
		switch c.State {
		case MainMenu:
			c.mainMenu()
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
