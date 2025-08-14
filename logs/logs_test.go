package logs

import (
	"testing"
)

func TestHtgm(t *testing.T) {
	day := "2025-07-01 "
	line := "09:00:18 [ToHost: POR [Pos: (11, 123, 4263, 1), AssigId: 06191855, TuType: 01, Load Info: 1, Dst: 111254268001, rCode: 000, TuIdP: 1, TuId: E045EE0004206203  ]]"
	var h = &HtgmMsg{}
	NewMsg(h, day + line)
}
