package llm

type LLM interface {
	Clone() LLM
	AddSystemPrompt(prompt string) error
	AddUserPrompt(prompt string) error
	Execute() error
}
