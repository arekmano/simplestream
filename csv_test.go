package simplestream_test

import (
	"encoding/csv"
	"io"
	"strings"
	"testing"

	"github.com/arekmano/simplestream"
	"github.com/stretchr/testify/require"
)

func TestCsvSource_valid(t *testing.T) {
	stringReader := strings.NewReader("Title,Note,URL,Comment")

	reader := simplestream.CsvSource{
		Reader: csv.NewReader(stringReader),
		RowDecode: func(row []string, p interface{}) error {
			require.Equal(t, row[0], "Title")
			require.Equal(t, row[1], "Note")
			require.Equal(t, row[2], "URL")
			require.Equal(t, row[3], "Comment")
			p.(*testStruct).Test = 20
			return nil
		},
	}

	output := testStruct{}
	err := reader.Fetch(&output)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, output.Test, 20)
}

func TestCsvSource_invalid(t *testing.T) {
	stringReader := strings.NewReader("")
	reader := simplestream.CsvSource{
		Reader: csv.NewReader(stringReader),
		RowDecode: func(row []string, p interface{}) error {
			return nil
		},
	}

	err := reader.Fetch(nil)
	if err != io.EOF {
		t.Fatal("Expected an EOF error")
	}
}
