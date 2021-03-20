package model

import (
	"encoding/json"

	"strings"
)

// 一个az对象
type AzInfo struct {
	// 可用分区的名字。

	ZoneName string `json:"zoneName"`

	ZoneState *ZoneState `json:"zoneState"`
}

func (o AzInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AzInfo struct{}"
	}

	return strings.Join([]string{"AzInfo", string(data)}, " ")
}
