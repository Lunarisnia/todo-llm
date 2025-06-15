package describe

import "github.com/Lunarisnia/todo-llm/internal/core/llm"

type DescribeAgent struct {
	Model llm.LLM
}

func NewDescribeAgent(model llm.LLM) *DescribeAgent {
	model.AddSystemPrompt(`
		Your task is to take the user input and analyze what the user might mean.
		Reply in a paragraph to describe what the user mean no need for any flavour text, be specific and don't omit anything the user mentioned
		`)
	return &DescribeAgent{
		Model: model,
	}
}

func (d *DescribeAgent) Chat(prompt string) (*llm.ChatResponse, error) {
	err := d.Model.AddUserPrompt(prompt)
	if err != nil {
		return nil, err
	}
	chatResponse, err := d.Model.Execute()
	if err != nil {
		return nil, err
	}

	return chatResponse, nil
}
