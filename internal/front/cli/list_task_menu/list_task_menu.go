package listtaskmenu

import (
	"fmt"

	"github.com/Lunarisnia/todo-llm/internal/core"
	"github.com/Lunarisnia/todo-llm/internal/core/task"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/clistate"
	"github.com/Lunarisnia/todo-llm/internal/front/cli/menu"
	"github.com/Lunarisnia/todo-llm/internal/input"
)

type ListTaskMenu struct {
	Options []menu.Option

	TodoEngine core.TodoEngine
}

func NewListTaskMenu(todoEngine core.TodoEngine) menu.Menu {
	listTaskMenu := &ListTaskMenu{
		Options:    make([]menu.Option, 0),
		TodoEngine: todoEngine,
	}

	listTaskMenu.Options = append(listTaskMenu.Options, menu.NewOption(
		"Your active task:",
		listTaskMenu.ListTask,
	))
	return listTaskMenu
}

func (l *ListTaskMenu) GetOptions() []menu.Option {
	return l.Options
}

func (l *ListTaskMenu) GetCursorIndex() int {
	return -1
}

func (l *ListTaskMenu) SetCursorIndex(index int) {

}

func (l *ListTaskMenu) Callback(index int) error {
	return l.Options[index].Action()
}

func (l *ListTaskMenu) ListTask() error {
	tasks, err := l.TodoEngine.ListTask()
	if err != nil {
		return err
	}

	for _, t := range tasks {
		l.DisplayTask(t, "-", 0)
	}

	input.Read("")
	clistate.State = clistate.MainMenu

	return nil
}

func (l *ListTaskMenu) DisplayTask(t *task.Task, prefix string, tabspace int) {
	fmt.Println(prefix, t.Content)
	if len(t.Subtask) > 0 {
		tabspace++
		prefix = ""
		for range tabspace {
			prefix += "\t"
		}
		prefix += "-"

		for _, subtask := range t.Subtask {
			l.DisplayTask(subtask, prefix, tabspace)
		}
	}
}
