package testTable

type Product struct {
	field int
}

func (s *Product) Process(f1 int) (string, error) {
	s.field = f1 // concurrent tests race to update
	return "Success", nil
}
