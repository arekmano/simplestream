package simplestream

type NoMoreSourceError struct {
}

func (n NoMoreSourceError) Error() string {
	return "No more errors found at source"
}
