package logs

import (
	"regexp"
	"time"

	"github.com/Lily34927/swisslog/utils"
)

type HstMsg struct {
	TimeStamp time.Time `json:"timestamp"`
	Reserved  string    `json:"reserved"` // 保留字段，不明含义
	Protocol  string    `json:"protocol"`
	Msg       string    `json:"msg"`
}

func (h *HstMsg) Parse(line string) error {
	re := regexp.MustCompile(`^(?P<timestamp>\d{4}-\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2})\s+(?P<reserved>\w+)\s+(?P<protocol>\w+)\s+(?P<msg>.*)$`)
	result, err := utils.GetRegExpMap(re, line)
	if err != nil {
		return err
	}

	timestamp, err := utils.StringToTime(result["timestamp"])
	if err != nil {
		return err
	}

	h.TimeStamp = timestamp
	h.Reserved = result["reserved"]
	h.Protocol = result["protocol"]
	h.Msg = result["msg"]
	return nil
}
