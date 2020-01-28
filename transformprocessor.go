package simplestream

import (
	"fmt"
	"io"
	"reflect"

	"github.com/pkg/errors"
)

type TransformProcessor struct {
	Source
	Transformer
	Destination
}

func (s *TransformProcessor) Process(input interface{}) error {
	if err := s.validate(); err != nil {
		return err
	}
	err := s.Source.Fetch(input)
	if err == (NoMoreSourceError{}) {
		return io.EOF
	}
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error occurred when fetching from source using %s", reflect.TypeOf(s.Source)))
	}
	output, err := s.Transform(input)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error occurred when transforming using %s", reflect.TypeOf(s.Transformer)))
	}
	err = s.Destination.Put(output)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error occurred when putting to the destination using %s", reflect.TypeOf(s.Destination)))
	}
	return nil
}

func (s *TransformProcessor) validate() error {
	if s.Source == nil {
		return errors.New("source is not defined")
	}
	if s.Destination == nil {
		return errors.New("destination is not defined")
	}
	if s.Transformer == nil {
		return errors.New("transformer is not defined")
	}
	return nil
}
