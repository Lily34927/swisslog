package protocols

import (
	"log"
	"testing"

	"github.com/Lily34927/swisslog/utils"
)

func TestARQ(t *testing.T) {
	msg := "Src: (11, 127, 4014, 1), Dst: (11, 141, 4015, 2), AssigId: 06191892, TuType: 1, IoData: 0, TuIdPresent: True, TuId: E045EE0003D0060B"
	var h = &HtgmARQ{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestACP(t *testing.T) {
	msg := "Dst: (11, 127, 4014, 1), AssigId: 06191892, ErrCode: 000"
	var h = &HtgmACP{}
	NewProtocols(h, msg)
}

func TestIPR(t *testing.T) {
	msg := "Pos: (21, 14, 4821, 1), AssigId: 00000000, TuType: 00, Height: 0, WidthRight: 0, WidthLeft: 0, LengthFront: 0, LengthBack: 0, Weight: 0000000, Tunnel: 0, Runner: 0, TuIdP: 1, TuId:                   , ScanErr: 000, LabelLength: 105, Label: 91018739113100012507210274144895"
	var h = &HtgmIPR{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestMSGOUT(t *testing.T) {
	msg := "MsgNumber: 004, Label: 210084888001111094232001N291018739113100012507070311719551;91018739113100012507070311719807;"
	var h = &HtgmMSGOUT{}
	NewProtocols(h, msg)
}

func TestMSG_(t *testing.T) {
	msg := "MsgNumber: 002, Label: 211044551001100000000000000000000"
	msgNumber, err := GetHtgmMSGNumber(msg)
	if err != nil {
		log.Panic("err错误：", err)
		return
	}

	var p Protocols
	switch msgNumber {
	case "002":
		p = &HtgmMSG002{}
		p.Parse(msg)
	default:
		log.Println("invalid number")
	}

	p.Parse(msg)
	log.Println(p)
}

func TestMSG001(t *testing.T) {
	msg := "MsgNumber: 001, Label: 01873921117460100100000110000000000000006"
	var h = &HtgmMSG001{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestMSG002(t *testing.T) {
	msg := "MsgNumber: 002, Label: 211044551001100000000000000000000"
	var h = &HtgmMSG002{}
	NewProtocols(h, msg)
}

func TestMSG003(t *testing.T) {
	msg := "MsgNumber: 003, Label: 91212236113100012509100032898199101"
	var h = &HtgmMSG003{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestMSG005(t *testing.T) {
	msg := "MsgNumber: 005, Label: 210024948001"
	var h = &HtgmMSG005{}
	NewProtocols(h, msg)
}

func TestMSG006(t *testing.T) {
	msg := "MsgNumber: 006, Label: 222024498001"
	var h = &HtgmMSG006{}
	NewProtocols(h, msg)
}

func TestMSG007(t *testing.T) {
	msg := "MsgNumber: 007, Label: 91075206113100012505270012577903102"
	var h = &HtgmMSG007{}
	NewProtocols(h, msg)
}

func TestMSG008(t *testing.T) {
	msg := "MsgNumber: 008, Label: 11105422200199999900"
	var h = &HtgmMSG008{}
	NewProtocols(h, msg)
}

func TestMSG009(t *testing.T) {
	msg := "MsgNumber: 009, Label: 0752061001"
	var h = &HtgmMSG009{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestMSG010(t *testing.T) {
	msg := "MsgNumber: 010, Label: 91075206113100012508020305281152000"
	var h = &HtgmMSG010{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestPOR(t *testing.T) {
	msg := "Pos: (11, 123, 4263, 1), AssigId: 06191855, TuType: 01, Load Info: 1, Dst: 111254268001, rCode: 000, TuIdP: 1, TuId: E045EE0004206203"
	var h = &HtgmPOR{}
	NewProtocols(h, msg)
}

func TestSDI(t *testing.T) {
	msg := "HostBitNo: 0003, Value: 1, ErrCode: 000"
	var h = &HtgmSDI{}
	NewProtocols(h, msg)
}

func TestSDO(t *testing.T) {
	msg := "HostBitId: 26, Value: 1"
	var h = &HtgmSDO{}
	NewProtocols(h, msg)
}

func TestHstARQ(t *testing.T) {
	msg := "05 06261379 CM 01 31-010-953-00-01 31-010-027-05-01 RE HI FU FU"
	var h = &HstARQ{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestHstACP(t *testing.T) {
	msg := "05 06261379 31-010-027-05-00 UL-UL UL-UL 000 0"
	var h = &HstACP{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}

func TestHstCSR(t *testing.T) {
	msg := "05 06261466 1 069486 UL-UL UL-UL 05 000"
	var h = &HstCSR{}
	NewProtocols(h, msg)
	utils.StructToMap(h)
}
