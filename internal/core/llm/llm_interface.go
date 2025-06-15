package llm

type LLM interface {
	AddSystemPrompt(prompt string) error
	AddUserPrompt(prompt string) error
	Execute() error
}
