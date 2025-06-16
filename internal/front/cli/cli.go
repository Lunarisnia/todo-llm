package cli

import (
	"fmt"
	"log"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front"
	mainmenu "github.com/Lunarisnia/todo-llm/internal/front/cli/main_menu"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/rendering"
)

type cliImpl struct {
	State      ProgramState
	TodoEngine core.TodoEngine

	Renderer rendering.RenderingEngine
	MainMenu menu.Menu
}

func NewCLIFrontend(todoEngine core.TodoEngine) front.FrontEnd {
	return &cliImpl{
		State:      MainMenu,
		TodoEngine: todoEngine,
		Renderer:   rendering.NewRenderingEngine(),
		MainMenu:   mainmenu.NewMainMenu(todoEngine),
	}
}

func (c *cliImpl) Run() {
	for {
		switch c.State {
		case MainMenu:
			// c.MainMenu.Render()
			c.Renderer.SetRenderObject(c.MainMenu)
		case ListAll:
			fmt.Println("ListAll")
		case AddNew:
			fmt.Println("AddNew")
		default:
			log.Fatalln("unknown or unimplemented state")
		}

		c.Renderer.Render()
	}
}

func (c *cliImpl) mainMenu() {
	fmt.Println("Main Menu")
}

func (c *cliImpl) SetState(newState ProgramState) {
	c.State = newState
}
