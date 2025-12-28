package page

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Home() Page {
	return Page{
		Title:       "Site Produto",
		Description: "Este é exatamente o site que você vai receber.",
	}
}
