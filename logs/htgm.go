package logs

import (
	"regexp"
	"time"

	"github.com/Lily34927/swisslog/utils"
)

type HtgmMsg struct {
	TimeStamp time.Time `json:"timestamp"`
	Direction string    `json:"direction"` // 方向：ToHost 或 FrHost
	Protocol  string    `json:"protocol"`
	Msg       string    `json:"msg"`
}

func (h *HtgmMsg) Parse(line string) error {
	re := regexp.MustCompile(`^(?P<timestamp>\d{4}-\d{2}-\d{2}\s*\d{2}:\d{2}:\d{2})\s+\[(?P<direction>\w+):\s*(?P<protocol>\w+)\s*\[(?P<msg>.*)\]\]$`)

	result, err := utils.GetRegExpMap(re, line)
	if err != nil {
		return err
	}

	timeStamp, err := utils.StringToTime(result["timestamp"])
	if err != nil {
		return err
	}

	h.TimeStamp = timeStamp
	h.Direction = result["direction"]
	h.Protocol = result["protocol"]
	h.Msg = result["msg"]

	return nil
}
