package model

import (
	"encoding/json"

	"strings"
)

type UpdateUnreadNewInstantMsgV2Req struct {
	// 组id

	GroupId *string `json:"group_id,omitempty"`
}

func (o UpdateUnreadNewInstantMsgV2Req) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUnreadNewInstantMsgV2Req struct{}"
	}

	return strings.Join([]string{"UpdateUnreadNewInstantMsgV2Req", string(data)}, " ")
}
