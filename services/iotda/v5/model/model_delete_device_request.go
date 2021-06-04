package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDeviceRequest struct {
	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// **参数说明**：设备ID，用于唯一标识一个设备。在注册设备时直接指定，或者由物联网平台分配获得。由物联网平台分配时，生成规则\"product_id\" + \"_\" + \"node_id\"拼接而成。 **取值范围**：长度不超过128，只允许字母、数字、下划线（_）、连接符（-）的组合。

	DeviceId string `json:"device_id"`
}

func (o DeleteDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeviceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceRequest", string(data)}, " ")
}
