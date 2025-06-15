package breakdown

import "github.com/Lunarisnia/todo-llm/internal/core/llm"

type breakDownImpl struct {
	Model llm.LLM
}

func NewBreakDownAgent(model llm.LLM) *breakDownImpl {
	agent := &breakDownImpl{
		Model: model,
	}

	agent.Model.AddSystemPrompt("You will always reply in a distinguished manner, treat the user like they are a member of the royalty and you are their servant")
	agent.Model.AddSystemPrompt("Only reply with the bullet point of what the user need to do, no need for any further explanation or flavour text")

	return agent
}

func (b *breakDownImpl) Chat(prompt string) error {
	err := b.Model.AddUserPrompt(prompt)
	if err != nil {
		return err
	}

	err = b.Model.Execute()
	if err != nil {
		return err
	}

	return nil
}
