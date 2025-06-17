package rendering

import (
	"fmt"
	"strconv"

	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

type RenderingEngine interface {
	Render() error
	SetRenderObject(object menu.Menu)
	SetCursorMode(mode bool)
}

var CursorMode bool = true

type renderingEngineImpl struct {
	RenderObject menu.Menu
}

func NewRenderingEngine() RenderingEngine {
	return &renderingEngineImpl{}
}

func (r *renderingEngineImpl) SetRenderObject(object menu.Menu) {
	r.RenderObject = object
}

func (r *renderingEngineImpl) Render() error {
	fmt.Print("\033[H\033[2J")
	for i, opt := range r.RenderObject.GetOptions() {
		if len(r.RenderObject.GetOptions()) == 1 {
			fmt.Println(opt.Title)
		} else {
			fmt.Println(i+1, opt.Title)
		}
	}
	if len(r.RenderObject.GetOptions()) == 1 {
		err := r.RenderObject.Callback(0)
		if err != nil {
			return err
		}
	} else {
		choice := input.Read("Enter your choice:")
		choiceIndex, err := strconv.Atoi(choice)
		if err != nil {
			return err
		}

		err = r.RenderObject.Callback(choiceIndex - 1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *renderingEngineImpl) SetCursorMode(mode bool) {
	CursorMode = mode
}
