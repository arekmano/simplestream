package simplestream

import (
	"encoding/json"
	"errors"
	"io"
)

type Json struct {
	io.Reader
	io.Writer
}

func (j *Json) Fetch(p interface{}) error {
	if j.Reader == nil {
		return errors.New("Writer is not Defined")
	}
	return json.NewDecoder(j.Reader).Decode(p)
}

func (j *Json) Put(p interface{}) error {
	if j.Writer == nil {
		return errors.New("Writer is not Defined")
	}
	encoder := json.NewEncoder(j.Writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(p)
}
