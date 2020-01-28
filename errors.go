package simplestream

type NoMoreSourceError struct {
}

func (n NoMoreSourceError) Error() string {
	return "No more entries found at source"
}
