package simplestream

type BasicProcessor struct {
	Source
	Destination
}

func (s *BasicProcessor) Process(input interface{}) error {
	err := s.Source.Fetch(input)
	if err != nil {
		return err
	}
	return s.Destination.Put(input)
}
