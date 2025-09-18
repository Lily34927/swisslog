package protocols

import "github.com/Lily34927/swisslog/utils"

// ARQ协议
type HstARQ struct {
	// 05 06261379 CM 01 31-010-953-00-01 31-010-027-05-01 RE HI FU FU
	CraneNumber   int    `json:"craneNumber"`
	AssignId      string `json:"assignId"`
	AssignType    string `json:"assignType"`
	TuType        string `json:"tuType"`
	Src           string `json:"src"`
	CranePosition string `json:"cranePosition"`
	Fork          string `json:"fork"`
	Speed         string `json:"speed"`
	RearFork      string `json:"rearFork"`
	FrontFork     string `json:"frontFork"`
}

func (h *HstARQ) Parse(msg string) error {
	result := utils.StringToSlices(msg)

	h.CraneNumber = utils.StringToInt(result[0])
	h.AssignId = result[1]
	h.AssignType = result[2]
	h.TuType = result[3]
	h.Src = utils.StringRemoveDashes(result[4])
	h.CranePosition = utils.StringRemoveDashes(result[5])
	h.Fork = result[6]
	h.Speed = result[7]
	h.RearFork = result[8]
	h.FrontFork = result[9]
	return nil
}

// ACP协议
type HstACP struct {
	CraneNumber     int    `json:"craneNumber"`
	AssignId        string `json:"assignId"`
	CranePosition   string `json:"cranePosition"` // 同ARQ中的CranePosition
	RearFork        string `json:"rearFork"`
	FrontFork       string `json:"frontFork"`
	ReturnCode      string `json:"returnCode"`
	InfoBlockNumber int    `json:"infoBlockNumber"`
}

func (h *HstACP) Parse(msg string) error {
	// 05 06261379 31-010-027-05-00 UL-UL UL-UL 000 0
	result := utils.StringToSlices(msg)

	h.CraneNumber = utils.StringToInt(result[0])
	h.AssignId = result[1]
	h.CranePosition = utils.StringRemoveDashes(result[2])
	h.RearFork = result[3]
	h.FrontFork = result[4]
	h.ReturnCode = result[5]
	h.InfoBlockNumber = utils.StringToInt(result[6])

	return nil
}

// CSR协议
type HstCSR struct {
	CraneNumber        int    `json:"craneNumber"`
	AssignId           string `json:"assignId"`
	CraneMode          string `json:"craneMode"`
	CraneAislePosition string `json:"craneAislePosition"`
	RearFork           string `json:"rearFork"`
	FrontFork          string `json:"frontFork"`
	CurrentAisle       string `json:"currentAisle"`
	ReturnCode         string `json:"returnCode"`
}

func (h *HstCSR) Parse(msg string) error {
	// 05 06261466 1 069486 UL-UL UL-UL 05 000
	result := utils.StringToSlices(msg)

	h.CraneNumber = utils.StringToInt(result[0])
	h.AssignId = result[1]
	h.CraneMode = result[2]
	h.CraneAislePosition = result[3]
	h.RearFork = result[4]
	h.FrontFork = result[5]
	h.CurrentAisle = result[6]
	h.ReturnCode = result[7]
	return nil
}
