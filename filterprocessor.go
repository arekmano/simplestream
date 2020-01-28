package simplestream

import (
	"fmt"
	"io"
	"reflect"

	"github.com/pkg/errors"
)

type FilterProcessor struct {
	Source
	FilterConditionFunc func(input interface{}) bool
	Destination
}

func (s *FilterProcessor) Process(input interface{}) error {
	err := s.Source.Fetch(input)
	if err == (NoMoreSourceError{}) {
		return io.EOF
	}
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error occured when fetching from source using %s", reflect.TypeOf(s.Source)))
	}
	if !s.FilterConditionFunc(input) {
		return nil
	}
	err = s.Destination.Put(input)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error occured when putting to the destination using %s", reflect.TypeOf(s.Destination)))
	}
	return nil
}
