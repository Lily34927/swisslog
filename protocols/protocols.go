package protocols

import (
	"swisslog_parse/utils"
)

type Protocols interface{
	Parse(string) error
}

func NewProtocols (p Protocols, msg string){
	p.Parse(msg)
	utils.StructToMap(p)
}