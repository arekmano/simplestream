package simplestream

type StringSource struct {
	Output string
}

func (s *StringSource) Fetch(p interface{}) error {
	return nil
}

type ErrorSource struct {
	Output error
}

func (s *ErrorSource) Fetch(p interface{}) error {
	return s.Output
}
