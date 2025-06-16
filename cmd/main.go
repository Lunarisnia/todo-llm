package main

import (
	"github.com/Lunarisnia/todo-llm/internal/core"
	lmstudio "github.com/Lunarisnia/todo-llm/internal/core/llm/lm_studio"
	"github.com/Lunarisnia/todo-llm/internal/front/cli"
)

func main() {
	llmModel := lmstudio.NewLMStudio(0.4, -1)
	todoEngine := core.NewTodoEngine(llmModel)

	frontend := cli.NewCLIFrontend(todoEngine)
	frontend.Run()
}
