package simplestream

import "github.com/pkg/errors"

type StringDestination struct {
	Output *string
}

func (s *StringDestination) Put(p interface{}) error {
	str, ok := p.(*string)
	if !ok {
		return errors.New("Error!")
	}
	s.Output = str
	return nil
}

type ErrorDestination struct {
	Output error
}

func (s *ErrorDestination) Put(p interface{}) error {
	return s.Output
}
