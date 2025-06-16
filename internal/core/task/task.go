package task

var NextAvailableID uint = 0

type Task struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	Subtask    *Task  `json:"subtask"`
	ParentTask *Task  `json:"parent_task"`
	Done       bool   `json:"done"`
}

func NewTask(content string, parentTask *Task, subtask *Task) Task {
	return Task{
		ID:         uint(NextAvailableID),
		Content:    content,
		ParentTask: parentTask,
		Subtask:    subtask,
		Done:       false,
	}
}

func (t *Task) SetSubtask(subtask *Task) {
	t.Subtask = subtask
}

func (t *Task) SetParentTask(parent *Task) {
	t.ParentTask = parent
}
