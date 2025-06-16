package menu

type Menu interface {
	GetCursorIndex() int
	SetCursorIndex(newIndex int)
	GetOptions() []Option
	Callback(index int) error
}
