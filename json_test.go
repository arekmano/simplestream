package simplestream_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/arekmano/simplestream"
)

type testStruct struct {
	Test int `json:"test"`
}

func TestJsonFetch_valid(t *testing.T) {
	// Test data
	j := simplestream.Json{
		Reader: strings.NewReader("{\"test\":1}"),
	}

	var tester testStruct
	// Execute
	err := j.Fetch(&tester)

	// Verify
	if err != nil {
		t.Fatal(err)
	}
}

func TestJsonFetch_invalid(t *testing.T) {
	// Test data
	j := simplestream.Json{
		Reader: strings.NewReader("\"test\":1}"),
	}

	var tester testStruct
	// Execute
	err := j.Fetch(&tester)

	// Verify
	if err == nil {
		t.Fatal("No error thrown")
	}
}

func TestJsonFetch_undefinedReader(t *testing.T) {
	// Test data
	j := simplestream.Json{}

	var tester testStruct
	// Execute
	err := j.Fetch(&tester)

	// Verify
	if err == nil {
		t.Fatal("No error thrown")
	}
}

func TestJsonPut_valid(t *testing.T) {
	// Test data
	j := simplestream.Json{
		Writer: bytes.NewBufferString(""),
	}

	tester := &testStruct{
		Test: 23,
	}

	// Execute
	err := j.Put(tester)

	// Verify
	if err != nil {
		t.Fatal(err)
	}
}

func TestJsonPut_undefinedWriter(t *testing.T) {
	// Test data
	j := simplestream.Json{}

	tester := &testStruct{
		Test: 23,
	}

	// Execute
	err := j.Put(tester)

	// Verify
	if err == nil {
		t.Fatal("No error thrown")
	}
}
