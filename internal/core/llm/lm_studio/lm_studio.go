package lmstudio

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/Lunarisnia/todo-llm/internal/core/http_client"
	"github.com/Lunarisnia/todo-llm/internal/core/llm"
)

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
	body, err := json.Marshal(l.llmInfo)
	if err != nil {
		return err
	}

	chatResponse := &ChatResponse{}
	// TODO: Put the url somewhere else later
	err = httpclient.Do.Post("http://192.168.0.154:1234/api/v0/chat/completions", body, &chatResponse)
	if err != nil {
		return err
	}
	for _, choice := range chatResponse.Choices {
		fmt.Println(choice.Message)
	}
	// TODO: Make this return the result instead of printing it
	return nil
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
