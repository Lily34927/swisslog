package logs

import "swisslog_parse/utils"

type Msg interface{
	Parse(string) error
}

func NewMsg(m Msg, line string){
	m.Parse(line)
	utils.StructToMap(m)
}