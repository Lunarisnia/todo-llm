package cli

type ProgramState string

const (
	MainMenu ProgramState = "MAIN_MENU"
	AddNew   ProgramState = "ADD_NEW"
	ListAll  ProgramState = "LIST_ALL"
)

func (e ProgramState) Value() string {
	return string(e)
}

func (e ProgramState) Compare(state ProgramState) bool {
	return e == state
}
