package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/clistate"
	mainmenu "github.com/Lunarisnia/todo-llm/internal/front/cli/main_menu"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/rendering"
)

type cliImpl struct {
	TodoEngine core.TodoEngine

	Renderer      rendering.RenderingEngine
	mainMenuScene menu.Menu
}

func NewCLIFrontend(todoEngine core.TodoEngine) front.FrontEnd {
	return &cliImpl{
		TodoEngine:    todoEngine,
		Renderer:      rendering.NewRenderingEngine(),
		mainMenuScene: mainmenu.NewMainMenu(todoEngine),
	}
}

func (c *cliImpl) Run() {
	for {
		switch clistate.State {
		case clistate.MainMenu:
			c.Renderer.SetRenderObject(c.mainMenuScene)
		case clistate.ListAll:
			fmt.Println("ListAll")
		case clistate.AddNew:
			fmt.Println("AddNew")
			os.Exit(1)
		default:
			log.Fatalln("unknown or unimplemented state")
		}

		c.Renderer.Render()
	}
}

func (c *cliImpl) mainMenu() {
	fmt.Println("Main Menu")
}
