package simplestream_test

import (
	"testing"

	"github.com/arekmano/simplestream"
	"github.com/pkg/errors"
)

func TestBasicProcessorProcess_valid(t *testing.T) {
	// Test data
	s := "sample"
	p := simplestream.BasicProcessor{
		Source:      &simplestream.StringSource{"Test"},
		Destination: &simplestream.StringDestination{&s},
	}
	var v string
	err := p.Process(&v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBasicProcessorProcess_invalid(t *testing.T) {
	// Test data
	s := "sample"
	p := simplestream.BasicProcessor{
		Source:      &simplestream.ErrorSource{errors.New("Error")},
		Destination: &simplestream.StringDestination{&s},
	}
	var v string
	err := p.Process(&v)
	if err == nil {
		t.Fatal("Error expected")
	}
}
