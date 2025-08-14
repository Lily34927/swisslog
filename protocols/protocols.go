package protocols

type Protocols interface {
	Parse(string) error
}

func NewProtocols(p Protocols, msg string) {
	p.Parse(msg)
}
