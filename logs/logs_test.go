package logs

import (
	"testing"

	"github.com/Lily34927/swisslog/utils"
)

func TestHtgm(t *testing.T) {
	// day := "2025-07-22 "
	line := "2025-07-22 00:20:23 [ToHost: MSG_OUT [MsgNumber: 004, Label: 210144828001111034216001N291018739113100012507210274067583;91018739113100012507210274097791;   ]]"
	var h = &HtgmMsg{}
	NewMsg(h, line)
	utils.StructToMap(h)
}

func TestHst(t *testing.T) {
	/*
		// 报文样例
			2025-07-19 00:05:51  7929 CSR 05 06261466 1 069486 UL-UL UL-UL 05 000 // CSR
			2025-07-19 00:06:38  7778 ARQ 05 06261379 CM 01 31-010-953-00-01 31-010-027-05-01 RE HI FU FU // ARQ
			2025-07-19 00:07:41  7934 ACP 05 06261379 31-010-027-05-00 UL-UL UL-UL 000 0 // ACP
	*/
	line := "2025-07-19 00:07:41  7934 ACP 05 06261379 31-010-027-05-00 UL-UL UL-UL 000 0"
	var h = &HstMsg{}
	NewMsg(h, line)
	utils.StructToMap(h)
}
