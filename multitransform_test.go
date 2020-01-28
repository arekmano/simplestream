package simplestream_test

import (
	"fmt"
	"github.com/arekmano/simplestream"
	"github.com/stretchr/testify/require"
	"testing"
)

type FailTransformer struct {
	*testing.T
}

func (m *FailTransformer) Transform(input interface{}) (interface{}, error) {
	m.T.Fatal("FailTransformer has run")
	return nil, nil
}

func testCases() []interface{} {
	return []interface{}{
		"",
		"test",
		1,
		23,
		nil,
	}
}
func TestMultitransform_noTransformers(t *testing.T) {
	// Test data
	trans := &simplestream.MultiTransformer{}
	for _, input := range testCases() {
		t.Run(fmt.Sprintf("input: %s", input), testError(t, trans, input))
	}
}

func testIdentity(t *testing.T, transformer simplestream.Transformer, input interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		// Execute
		output, err := transformer.Transform(input)

		// Verify
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, input, output)
	}
}

func TestMultitransform_singleTransformer(t *testing.T) {
	// Test data
	trans := &simplestream.MultiTransformer{
		Transformers: []simplestream.Transformer{
			&simplestream.IdentityTransformer{},
		},
	}
	for _, input := range testCases() {
		t.Run(fmt.Sprintf("input: %s", input), testIdentity(t, trans, input))
	}
}

func TestMultitransform_multipleTransformer(t *testing.T) {
	// Test data
	trans := &simplestream.MultiTransformer{
		Transformers: []simplestream.Transformer{
			&simplestream.IdentityTransformer{},
			&simplestream.IdentityTransformer{},
		},
	}
	for _, input := range testCases() {
		t.Run(fmt.Sprintf("input: %s", input), testIdentity(t, trans, input))
	}
}

func TestMultitransform_errorTransformer(t *testing.T) {
	// Test data
	trans := &simplestream.MultiTransformer{
		Transformers: []simplestream.Transformer{
			&simplestream.ErrorTransformer{},
			&FailTransformer{T: t},
		},
	}
	for _, input := range testCases() {
		t.Run(fmt.Sprintf("input: %s", input), testError(t, trans, input))
	}
}

func testError(t *testing.T, transformer simplestream.Transformer, input interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		// Execute
		output, err := transformer.Transform(input)

		// Verify
		if err == nil {
			t.Fatal("Error was not thrown!")
		}
		require.Nil(t, output)
	}
}
