package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type TagDeviceRequest struct {
	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *BindTagsDto `json:"body,omitempty"`
}

func (o TagDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TagDeviceRequest struct{}"
	}

	return strings.Join([]string{"TagDeviceRequest", string(data)}, " ")
}
