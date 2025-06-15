package lmstudio

import "github.com/Lunarisnia/todo-llm/internal/core/llm"

type lmStudioImpl struct {
	llmInfo LLMInfo
}

func NewLMStudio(temperature float32, maxToken int) llm.LLM {
	return &lmStudioImpl{
		llmInfo: LLMInfo{
			Messages:    make([]LMStudioMessage, 0),
			Temperature: temperature,
			MaxToken:    maxToken,
		},
	}
}

func (l *lmStudioImpl) AddSystemPrompt(prompt string) error {
	message := LMStudioMessage{
		Role:    LMStudioSystem,
		Content: prompt,
	}
	l.llmInfo.AddMessage(message)
	return nil
}

func (l *lmStudioImpl) AddUserPrompt(prompt string) error {
	message := LMStudioMessage{
		Role:    LMStudioUser,
		Content: prompt,
	}
	l.llmInfo.AddMessage(message)
	return nil
}

func (l *lmStudioImpl) Execute() error {
	// TODO: Create http abstraction
	return nil
}
