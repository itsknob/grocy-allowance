package tui

type Service struct {
	DepositModel
	modelModalUi
}

func (s *Service) Name() string {
	return "Modal TUI"
}
