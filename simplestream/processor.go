package simplestream

import "github.com/pkg/errors"

import "fmt"

import "reflect"

import "io"

type BasicProcessor struct {
	Source
	Destination
}

func (s *BasicProcessor) Process(input interface{}) error {
	err := s.Source.Fetch(input)
	if err != nil {
		return err
	}
	return s.Destination.Put(input)
}

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
		return errors.Wrap(err, fmt.Sprintf("Error occured when fetching from source using %s", reflect.TypeOf(s.Source)))
	}
	output, err := s.Transform(input)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error occured when transforming using %s", reflect.TypeOf(s.Transformer)))
	}
	err = s.Destination.Put(output)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error occured when putting to the destination using %s", reflect.TypeOf(s.Destination)))
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
