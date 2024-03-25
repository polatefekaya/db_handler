package usecases

type AutomationUseCase struct {
	Process *Process
}

func NewAutomationUseCase() *AutomationUseCase {
	return &AutomationUseCase{}
}

func (m *AutomationUseCase) AutomatePlayer() {
	m.Process.Start()
}
