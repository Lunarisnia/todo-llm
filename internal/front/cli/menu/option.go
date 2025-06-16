package menu

type OptionAction func() error

type Option struct {
	Title  string
	Action OptionAction
}

func NewOption(title string, action OptionAction) Option {
	return Option{
		Title:  title,
		Action: action,
	}
}
