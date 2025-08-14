package logs

import (
	"testing"
)

func TestHtgm(t *testing.T) {
	// day := "2025-07-22 "
	line := "2025-07-22 00:20:23 [ToHost: MSG_OUT [MsgNumber: 004, Label: 210144828001111034216001N291018739113100012507210274067583;91018739113100012507210274097791;   ]]"
	var h = &HtgmMsg{}
	NewMsg(h, line)
}
