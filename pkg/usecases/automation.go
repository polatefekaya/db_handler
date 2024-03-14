package usecases

type AutomationUseCase struct {
	FootballUsecase *FootballUsecase
}

func NewAutomationUseCase() *AutomationUseCase {
	return &AutomationUseCase{}
}

func (m *AutomationUseCase) AutomatePlayer() {

}
