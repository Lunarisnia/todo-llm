package task

type Task struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	Subtask    *Task  `json:"subtask"`
	ParentTask *Task  `json:"parent_task"`
}

func (t *Task) SetSubtask(subtask *Task) {
	t.Subtask = subtask
}

func (t *Task) SetParentTask(parent *Task) {
	t.ParentTask = parent
}
