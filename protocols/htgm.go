package protocols

import (
	"fmt"
	"strings"

	"github.com/Lily34927/swisslog/utils"
)

// ARQ协议
type HtgmARQ struct {
	Src         string `json:"src"`
	Dst         string `json:"dst"`
	AssigId     string `json:"assigId"`
	TuType      string `json:"tuType"`
	IoData      int    `json:"ioData"`
	TuIdPresent bool   `json:"tuIdPresent"`
	TuId        string `json:"tuId"`
}

func (h *HtgmARQ) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	h.Src = utils.ParsePosition(results["Src"])
	h.Dst = utils.ParsePosition(results["Dst"])
	h.AssigId = results["AssigId"]
	h.TuType = results["TuType"]
	h.IoData = utils.StringToInt(results["IoData"])
	h.TuIdPresent = utils.StringToBool(results["TuIdPresent"])
	h.TuId = results["TuId"]
	return nil
}

// ACP协议
type HtgmACP struct {
	Dst     string `json:"dst"`
	AssigId string `json:"assigId"`
	ErrCode string `json:"errCode"`
}

func (h *HtgmACP) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	h.Dst = utils.ParsePosition(results["Dst"])
	h.AssigId = results["AssigId"]
	h.ErrCode = results["ErrCode"]
	return nil
}

// IPR协议
type HtgmIPR struct {
	Pos         string `json:"position"`    // 工位
	AssigId     string `json:"assigId"`     // 事务ID
	TuType      string `json:"tyType"`      // 类型
	Height      string `json:"height"`      // 件烟箱的基础信息，高度
	WidthRight  string `json:"widthRight"`  // 件烟箱的基础信息，右宽度
	WidthLeft   string `json:"widthLeft"`   // 件烟箱的基础信息，左宽度
	LengthFront string `json:"lengthFront"` // 件烟箱的基础信息，前长度
	LengthBack  string `json:"lengthBack"`  // 件烟箱的基础信息，后长度
	Weight      string `json:"weight"`      // 件烟箱的基础信息，重量
	Tunnel      string `json:"tunnel"`      // 隧道标识
	Runner      string `json:"runner"`      // 运行标识
	TuIdP       string `json:"tuIdP"`       // ID标识
	TuId        string `json:"tuId"`        // ID
	ScanErr     string `json:"scanErr"`     // 扫描错误码
	LabelLength int    `json:"labelLength"` // 标签长度，不准确
	Label       string `json:"label"`       // 标签内容
}

func (h *HtgmIPR) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	h.Pos = utils.ParsePosition(results["Pos"])
	h.AssigId = results["AssigId"]
	h.TuType = results["TuType"]
	h.Height = results["Height"]
	h.WidthRight = results["WidthRight"]
	h.WidthLeft = results["WidthLeft"]
	h.LengthFront = results["LengthFront"]
	h.LengthBack = results["LengthBack"]
	h.Weight = results["Weight"]
	h.Tunnel = results["Tunnel"]
	h.Runner = results["Runner"]
	h.TuIdP = results["TuIdP"]
	h.TuId = results["TuId"]
	h.ScanErr = results["ScanErr"] // 设置对应为名称
	h.LabelLength = utils.StringToInt(results["LabelLength"])
	h.Label = results["Label"]
	return nil
}

// MSGOUT协议
type HtgmMSGOUT struct {
	MsgNumber string   `json:"msgNumber"`
	Pos       string   `json:"position"` // 1~12位：码垛通道末尾工位
	Dst       string   `json:"dst"`      // 13~24位：对应的目标码垛工位
	IsLast    bool     `json:"isLast"`   // 25位：是否最后一次码垛
	BoxCount  int      `json:"boxCount"` // 26位：本次码垛的箱的数量(1~3)
	BoxIDs    []string `json:"boxIDs"`   // 件烟箱ID清单(32位)，最多3个，最少1个，用分号分隔
}

func (h *HtgmMSGOUT) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	label := results["Label"]
	h.MsgNumber = results["MsgNumber"]
	h.Pos = label[0:12]
	h.Dst = label[12:24]
	h.IsLast = utils.StringToBool(label[24:25])
	h.BoxCount = utils.StringToInt(label[25:26])
	h.BoxIDs = strings.Split(label[26:], ";")
	return nil
}

// MSG协议
type HtgmMSG interface { // 和Protocols重复，为了方便理解，供 MSG 协议调用
	Parse(msg string) error
}

type MsgHeader struct {
	MsgNumber string `json:"msgNumber"`
}

func NewHtgmMSG(msg string) (HtgmMSG, error) {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return nil, err
	}

	msgNumber := results["MsgNumber"]
	var h = MsgHeader{MsgNumber: msgNumber}
	switch msgNumber {
	case "001":
		return &HtgmMSG001{MsgHeader: h}, nil
	case "002":
		return &HtgmMSG002{MsgHeader: h}, nil
	case "005":
		return &HtgmMSG005{MsgHeader: h}, nil
	case "006":
		return &HtgmMSG006{MsgHeader: h}, nil
	case "007":
		return &HtgmMSG007{MsgHeader: h}, nil
	case "008":
		return &HtgmMSG008{MsgHeader: h}, nil
	default:
		return nil, fmt.Errorf("不在解析范围内, msg number: %s", msgNumber)
	}
}

// MSG001协议
type HtgmMSG001 struct {
	MsgHeader
	CigaretteCode string `json:"cigaretteCode"` // 1~6位：品牌规格代码的后六位(卷烟牌号)
	Src           string `json:"src"`           // 7~18位：工位码，Packaging Line(包装线)的工位
	LaneNumber    int    `json:"laneNumber"`    // 19-38位: 巷道号(物理位置从左到右)
	StackType     int    `json:"stackType"`     // 39-41位: 码垛类型
}

func (h *HtgmMSG001) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	label := results["Label"]
	h.CigaretteCode = label[0:6]
	h.Src = label[6:18]
	h.LaneNumber = utils.GetLaneNumber(label[18:38])
	h.StackType = utils.StringToInt(label[38:41])
	return nil
}

// MSG002协议
type HtgmMSG002 struct {
	MsgHeader
	Src        string `json:"src"`        // 1~12位：工位码，包装线(packaging line)工位
	Status     int    `json:"status"`     // 13位：订单状态，1代表订单结束(Closed), 2代表订单完成(Completed)
	LaneNumber int    `json:"laneNumber"` // 19-38位: 巷道号(物理位置从左到右)
}

func (h *HtgmMSG002) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	label := results["Label"]
	h.Src = label[0:12]
	h.Status = utils.StringToInt(label[12:13])
	h.LaneNumber = utils.GetLaneNumber(label[13:23])
	return nil
}

// MSG005协议
type HtgmMSG005 struct {
	MsgHeader
	Src string `json:"src"` // 1~12位：码垛巷道末尾(抓取)工位
}

func (h *HtgmMSG005) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	label := results["Label"]
	h.Src = label[0:12]
	return nil
}

// MSG006协议
type HtgmMSG006 struct {
	MsgHeader
	Src string `json:"src"` // 1~12位：注册装箱线工位
}

func (h *HtgmMSG006) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	label := results["Label"]
	h.Src = label[0:12]
	return nil
}

// MSG007协议
type HtgmMSG007 struct {
	MsgHeader
	BarCodeID string `json:"barCodeID"` // 1~32位：件烟箱的32位ID号
	ErrCode   string `json:"errCode"`   // 33~35位：WMS返回的错误码(Error Code)，订单已经关闭（101），无效的件烟箱ID（102），无效的规则ID（103）
}

func (h *HtgmMSG007) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	label := results["Label"]
	h.BarCodeID = label[0:32]
	h.ErrCode = label[32:35]
	return nil
}

// MSG008协议
type HtgmMSG008 struct {
	MsgHeader
	Dst        string `json:"dst"`        // 1~12位：工位码，码垛工位
	PalletType string `json:"palletType"` // 13~18位：托盘类型
	BoxCount   int    `json:"boxCount"`   // 19~20位：指出托盘上有几箱件烟，通常是00
}

func (h *HtgmMSG008) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	label := results["Label"]
	h.Dst = label[0:12]
	h.PalletType = label[12:18]
	h.BoxCount = utils.StringToInt(label[18:20])
	return nil
}

// POR协议
type HtgmPOR struct {
	Pos      string `json:"position"` // 工位
	AssigId  string `json:"assigId"`
	TuType   string `json:"tuType"`
	LoadInfo string `json:"loadInfo"` // 注意
	Dst      string `json:"dst"`
	RCode    string `json:"RCode"` // 注意
	TuIdP    string `json:"tuIdP"` // ID标识
	TuId     string `json:"tuId"`  // ID
}

func (h *HtgmPOR) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	h.Pos = utils.ParsePosition(results["Pos"])
	h.AssigId = results["AssigId"]
	h.TuType = results["TuType"]
	h.LoadInfo = results["Load Info"]
	h.Dst = results["Dst"]
	h.RCode = results["rCode"]
	h.TuIdP = results["TuIdP"]
	h.TuId = results["TuId"]
	return nil
}

// SDI协议
type HtgmSDI struct {
	HostBitNo string `json:"hostBitNo"`
	Value     int    `json:"value"`
	ErrCode   string `json:"errCode"`
}

func (h *HtgmSDI) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	h.HostBitNo = results["HostBitNo"]
	h.Value = utils.StringToInt(results["Value"])
	h.ErrCode = results["ErrCode"]
	return nil
}

// SDO协议
type HtgmSDO struct {
	HostBitId string `json:"hostBitId"`
	Value     int    `json:"value"`
}

func (h *HtgmSDO) Parse(msg string) error {
	results, err := utils.ParseMsg(msg)
	if err != nil {
		return err
	}

	h.HostBitId = results["HostBitId"]
	h.Value = utils.StringToInt(results["Value"])
	return nil
}
