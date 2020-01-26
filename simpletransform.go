package simplestream

import "errors"

type IdentityTransformer struct {
	Transformers []Transformer
}

func (m *IdentityTransformer) Transform(input interface{}) (interface{}, error) {
	return input, nil
}

type ErrorTransformer struct{}

func (m *ErrorTransformer) Transform(input interface{}) (interface{}, error) {
	return nil, errors.New("Using Error Transformer")
}
