package simplestream

import (
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"io"
)

type CsvSource struct {
	Reader    *csv.Reader
	RowDecode func(row []string, p interface{}) error
}

func (c *CsvSource) Fetch(p interface{}) error {
	row, err := c.Reader.Read()
	if err == io.EOF {
		return NoMoreSourceError{}
	}
	if err != nil {
		return err
	}
	logrus.
		WithField("row", row).
		Debug("Decoding Row")
	return c.RowDecode(row, p)
}
