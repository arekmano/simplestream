package simplestream

type Destination interface {
	Put(p interface{}) error
}

type Source interface {
	Fetch(p interface{}) error
}

type Processor interface {
	Process(input interface{})
}

type Transformer interface {
	Transform(input interface{}) (interface{}, error)
}
