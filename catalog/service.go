package catalog

type Service struct {
	Name string
	Rate int
	Sessions int
}

func (s Service) Total() int {
	return s.Rate * s.Sessions
}