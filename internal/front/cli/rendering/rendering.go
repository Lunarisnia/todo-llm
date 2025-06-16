package rendering

import (
	"fmt"

	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

type RenderingEngine interface {
	Render()
	SetRenderObject(object menu.Menu)
}

type renderingEngineImpl struct {
	RenderObject menu.Menu
}

func NewRenderingEngine() RenderingEngine {
	return &renderingEngineImpl{}
}

func (r *renderingEngineImpl) SetRenderObject(object menu.Menu) {
	r.RenderObject = object
}

func (r *renderingEngineImpl) Render() {
	cursorIndex := r.RenderObject.GetCursorIndex()
	fmt.Print("\033[H\033[2J")
	for i, opt := range r.RenderObject.GetOptions() {
		prefix := ""
		if cursorIndex == i {
			prefix = "*"
		}
		fmt.Println(prefix, opt.Title)
	}
	char, _ := input.ReadKeyboard("")
	if char == 'j' {
		cursorIndex++
	}
	if char == 'k' {
		cursorIndex--
		if cursorIndex < 0 {
			cursorIndex = len(r.RenderObject.GetOptions()) - 1
		}
	}
	r.RenderObject.SetCursorIndex(cursorIndex % len(r.RenderObject.GetOptions()))

}
