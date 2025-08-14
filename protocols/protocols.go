package protocols

import (
	"github.com/Lily34927/swisslog/utils"
)

type Protocols interface {
	Parse(string) error
}

func NewProtocols(p Protocols, msg string) {
	p.Parse(msg)
	utils.StructToMap(p)
}
