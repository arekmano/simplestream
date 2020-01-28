package simplestream

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type MultiTransformer struct {
	Transformers []Transformer
}

func (m *MultiTransformer) Transform(input interface{}) (interface{}, error) {
	o := input
	if m.Transformers == nil {
		return nil, errors.New("transformers are not defined")
	}
	for i, t := range m.Transformers {
		logrus.
			WithField("transformer", reflect.TypeOf(t)).
			Debug("Starting Transformer #", i)
		newO, err := t.Transform(o)
		if err != nil {
			return nil, err
		}
		o = newO
	}
	return o, nil
}
