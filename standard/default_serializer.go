package standard

type Serializer interface {
	Serialize() ([]byte, error)
}

type Unserializer interface {
	Unserialize([]byte) (interface{}, uint64, error)
}

type Namer interface {
	Name() string
}