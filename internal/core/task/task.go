package task

var NextAvailableID uint = 0

type Task struct {
	ID         uint    `json:"id"`
	Content    string  `json:"content"`
	Subtask    []*Task `json:"subtask"`
	ParentTask *Task   `json:"parent_task"`
	Done       bool    `json:"done"`
}

func NewTask(content string, parentTask *Task) Task {
	return Task{
		ID:         uint(NextAvailableID),
		Content:    content,
		ParentTask: parentTask,
		Subtask:    make([]*Task, 0),
		Done:       false,
	}
}

func (t *Task) AddSubtask(subtask *Task) {
	subtask.SetParentTask(t)
	t.Subtask = append(t.Subtask, subtask)
}

func (t *Task) SetParentTask(parent *Task) {
	t.ParentTask = parent
}
