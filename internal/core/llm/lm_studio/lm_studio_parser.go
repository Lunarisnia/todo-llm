package lmstudio

import (
	"strings"

	"github.com/Lunarisnia/todo-llm/internal/core/llm"
)

func ParseBulletPoint(chatResponse *llm.ChatResponse) []string {
	message := chatResponse.Choices[0].Message.Content

	tasks := strings.Split(message, "\n")
	for i := range tasks {
		tasks[i] = strings.TrimPrefix(tasks[i], "*")
		tasks[i] = strings.TrimSpace(tasks[i])
	}

	return tasks
}
