package lmstudio

type LMStudioRole string

const (
	LMStudioSystem LMStudioRole = "system"
	LMStudioUser   LMStudioRole = "user"
)

func (l LMStudioRole) String() string {
	return string(l)
}

type LMStudioMessage struct {
	Role    LMStudioRole `json:"role"`
	Content string       `json:"content"`
}

type LLMInfo struct {
	Messages    []LMStudioMessage `json:"messages"`
	Temperature float32           `json:"temperature"`
	MaxToken    int               `json:"max_tokens"`
}

func (l *LLMInfo) AddMessage(message LMStudioMessage) {
	l.Messages = append(l.Messages, message)
}
