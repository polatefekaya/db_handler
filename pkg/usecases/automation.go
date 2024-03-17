package usecases

type AutomationUseCase struct {
	FootballUsecase *FootballUsecase
	Process         *Process
}

func NewAutomationUseCase() *AutomationUseCase {
	return &AutomationUseCase{}
}

func (m *AutomationUseCase) AutomatePlayer() {
	m.Process.Start()
}
