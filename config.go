package shortURLz

type Config struct {
	Token   string
	RunMode Mode
}

type Mode int

const (
	Shorten Mode = iota + 1
	List
	ListGroup
	Delete
)

func NewConfig(token string, mode Mode) *Config {
	return &Config{Token: token, RunMode: mode}
}

func (m Mode) String() string {
	switch m {
	case Shorten:
		return "shorten"
	case List:
		return "list"
	case ListGroup:
		return "listgroup"
	case Delete:
		return "delete"
	default:
		return "unknown"
	}
}
