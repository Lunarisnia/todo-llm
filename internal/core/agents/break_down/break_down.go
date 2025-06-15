package breakdown

import (
	"github.com/Lunarisnia/todo-llm/internal/core/llm"
)

type BreakDownAgent struct {
	Model llm.LLM
}

func NewBreakDownAgent(model llm.LLM) *BreakDownAgent {
	agent := &BreakDownAgent{
		Model: model,
	}

	agent.Model.AddSystemPrompt(`
		Your task is to break down task that is described by the user,
		Always start your bullet point with the character '*' and no other formatting characters,
		Only reply with the bullet point of what the user need to do, no need for any further explanation or flavour text.
		`)

	return agent
}

func (b *BreakDownAgent) Chat(prompt string) (*llm.ChatResponse, error) {
	err := b.Model.AddUserPrompt(prompt)
	if err != nil {
		return nil, err
	}

	chatResponse, err := b.Model.Execute()
	if err != nil {
		return nil, err
	}

	return chatResponse, nil
}
