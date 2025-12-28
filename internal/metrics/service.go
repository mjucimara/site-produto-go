package metrics

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) All() []Metric {
	return []Metric{
		{
			Label: "Tempo médio de carregamento",
			Value: "~250 ms",
			Note:  "Medido localmente com hey, 1000 requisições, sem cache.",
		},
		{
			Label: "Peso total da página",
			Value: "< 50 KB",
			Note:  "HTML e CSS sem frameworks ou assets pesados.",
		},
		{
			Label: "Número de requisições",
			Value: "1",
			Note:  "HTML entregue em uma única resposta.",
		},
	}
}
