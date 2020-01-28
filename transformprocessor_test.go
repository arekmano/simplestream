package simplestream_test

import (
	"testing"

	"github.com/arekmano/simplestream"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestTransformProcessorProcess_valid(t *testing.T) {
	// Test data
	s := "sample"
	p := simplestream.TransformProcessor{
		Source:      &simplestream.StringSource{"Test"},
		Transformer: &simplestream.IdentityTransformer{},
		Destination: &simplestream.StringDestination{&s},
	}
	var v string
	err := p.Process(&v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransformProcessorProcess_missingTransformer(t *testing.T) {
	// Test data
	s := "sample"
	p := simplestream.TransformProcessor{
		Source:      &simplestream.ErrorSource{errors.New("Error")},
		Destination: &simplestream.StringDestination{&s},
	}
	var v string
	err := p.Process(&v)
	if err == nil {
		t.Fatal("Error expected")
	}
	require.Equal(t, "transformer is not defined", err.Error())
}

func TestTransformProcessorProcess_missingDestination(t *testing.T) {
	// Test data
	p := simplestream.TransformProcessor{
		Source:      &simplestream.ErrorSource{errors.New("Error")},
		Transformer: &simplestream.IdentityTransformer{},
	}
	var v string
	err := p.Process(&v)
	if err == nil {
		t.Fatal("Error expected")
	}
	require.Equal(t, "destination is not defined", err.Error())
}

func TestTransformProcessorProcess_missingSource(t *testing.T) {
	// Test data
	s := "sample"
	p := simplestream.TransformProcessor{
		Transformer: &simplestream.IdentityTransformer{},
		Destination: &simplestream.StringDestination{&s},
	}
	var v string
	err := p.Process(&v)
	if err == nil {
		t.Fatal("Error expected")
	}
	require.Equal(t, "source is not defined", err.Error())
}

func TestTransformProcessorProcess_invalidSource(t *testing.T) {
	// Test data
	s := "sample"
	p := simplestream.TransformProcessor{
		Source:      &simplestream.ErrorSource{errors.New("error")},
		Transformer: &simplestream.IdentityTransformer{},
		Destination: &simplestream.StringDestination{&s},
	}
	var v string
	err := p.Process(&v)
	if err == nil {
		t.Fatal("Error expected")
	}
}

func TestTransformProcessorProcess_invalidTransform(t *testing.T) {
	// Test data
	s := "sample"
	p := simplestream.TransformProcessor{
		Source:      &simplestream.StringSource{"Test"},
		Transformer: &simplestream.ErrorTransformer{},
		Destination: &simplestream.StringDestination{&s},
	}
	var v string
	err := p.Process(&v)
	if err == nil {
		t.Fatal("Error expected")
	}
}

func TestTransformProcessorProcess_invalidDestination(t *testing.T) {
	// Test data
	p := simplestream.TransformProcessor{
		Source:      &simplestream.StringSource{"Test"},
		Transformer: &simplestream.IdentityTransformer{},
		Destination: &simplestream.ErrorDestination{errors.New("error")},
	}
	var v string
	err := p.Process(&v)
	if err == nil {
		t.Fatal("Error expected")
	}
}
