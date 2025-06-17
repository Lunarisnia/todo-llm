package cli

import (
	"fmt"
	"log"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/front"
	addnewmenu "github.com/Lunarisnia/todo-llm/internal/front/cli/add_new_menu"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/clistate"
	mainmenu "github.com/Lunarisnia/todo-llm/internal/front/cli/main_menu"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/rendering"
)

type cliImpl struct {
	TodoEngine core.TodoEngine

	Renderer        rendering.RenderingEngine
	mainMenuScene   menu.Menu
	addNewMenuScene menu.Menu
}

func NewCLIFrontend(todoEngine core.TodoEngine) front.FrontEnd {
	return &cliImpl{
		TodoEngine:      todoEngine,
		Renderer:        rendering.NewRenderingEngine(),
		mainMenuScene:   mainmenu.NewMainMenu(todoEngine),
		addNewMenuScene: addnewmenu.NewAddNewMenu(todoEngine),
	}
}

func (c *cliImpl) Run() {
	for {
		switch clistate.State {
		case clistate.MainMenu:
			c.Renderer.SetRenderObject(c.mainMenuScene)
		case clistate.ListAll:
			// TODO: Implement list all
			fmt.Println("ListAll")
		case clistate.AddNew:
			c.Renderer.SetRenderObject(c.addNewMenuScene)
		default:
			log.Fatalln("unknown or unimplemented state")
		}

		c.Renderer.Render()
	}
}
