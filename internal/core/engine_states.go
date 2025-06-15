package core

type EngineState string

const (
	MainMenu EngineState = "MAIN_MENU"
	AddNew   EngineState = "ADD_NEW"
	ListAll  EngineState = "LIST_ALL"
)

func (e EngineState) Value() string {
	return string(e)
}

func (e EngineState) Compare(state EngineState) bool {
	return e == state
}
