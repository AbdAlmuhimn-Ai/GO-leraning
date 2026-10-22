package rental

type Service struct {
	Name     string
	Fee      int
	Sessions int
}

func (s Service) Value() int {
	return s.Fee * s.Sessions
}
