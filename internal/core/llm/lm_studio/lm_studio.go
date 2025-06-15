package lmstudio

import (
	"encoding/json"

	httpclient "github.com/Lunarisnia/todo-llm/internal/core/http_client"
	"github.com/Lunarisnia/todo-llm/internal/core/llm"
)

type lmStudioImpl struct {
	userPrompt LMStudioMessage
	llmInfo    LLMInfo
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
	l.userPrompt = LMStudioMessage{
		Role:    LMStudioUser,
		Content: prompt,
	}
	return nil
}

func (l *lmStudioImpl) Execute() (*llm.ChatResponse, error) {
	info := l.llmInfo
	info.Messages = make([]LMStudioMessage, 0)
	for _, message := range l.llmInfo.Messages {
		info.Messages = append(info.Messages, message)
	}
	info.Messages = append(info.Messages, l.userPrompt)

	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	chatResponse := llm.ChatResponse{}
	// TODO: Put the url somewhere else later
	err = httpclient.Do.Post("http://192.168.0.154:1234/api/v0/chat/completions", body, &chatResponse)
	if err != nil {
		return nil, err
	}

	return &chatResponse, nil
}

func (l *lmStudioImpl) Clone() llm.LLM {
	return &lmStudioImpl{
		llmInfo: LLMInfo{
			Messages:    make([]LMStudioMessage, 0),
			Temperature: l.llmInfo.Temperature,
			MaxToken:    l.llmInfo.MaxToken,
		},
	}
}
