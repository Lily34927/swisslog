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
	h, err := NewHtgmMSG(msg)
	if err != nil {
		log.Panic("err错误：", err)
		return
	}
	NewProtocols(h, msg)

	h2, ok := h.(*HtgmMSG002)
	if ok {
		log.Println("Parsed MsgNumber from HtgmMSG002:", h2.MsgNumber)
	} else {
		log.Println("h is not of type *HtgmMSG002")
	}

	// 如果你已经确认 h 是 *HtgmMSG002 类型，打印 MsgNumber
	log.Println("Final MsgNumber:", h2.MsgNumber)

	// switch v := h.(type) {
	// case *HtgmMSG001:
	// 	h = v
	// case *HtgmMSG002:
	// 	h = v
	// case *HtgmMSG005:
	// 	log.Println("MsgNumber from HtgmMSG005 after Parse:", v.MsgNumber)
	// case *HtgmMSG006:
	// 	log.Println("MsgNumber from HtgmMSG006 after Parse:", v.MsgNumber)
	// case *HtgmMSG007:
	// 	log.Println("MsgNumber from HtgmMSG007 after Parse:", v.MsgNumber)
	// case *HtgmMSG008:
	// 	log.Println("MsgNumber from HtgmMSG008 after Parse:", v.MsgNumber)
	// default:
	// 	log.Println("Unknown type")
	// }
	// log.Println(h.MsgNumber)

}

func TestMSG001(t *testing.T) {
	msg := "MsgNumber: 001, Label: 07520621106453100100001000000000000000001"
	var h = &HtgmMSG001{}
	NewProtocols(h, msg)
}

func TestMSG002(t *testing.T) {
	msg := "MsgNumber: 002, Label: 211044551001100000000000000000000"
	var h = &HtgmMSG002{}
	NewProtocols(h, msg)
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
