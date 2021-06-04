package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDevicesInGroupRequest struct {
	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// **参数说明**：设备组ID，用于唯一标识一个设备组，在创建设备组时由物联网平台分配。 **取值范围**：长度不超过36，十六进制字符串和连接符（-）的组合。

	GroupId string `json:"group_id"`
	// **参数说明**：分页查询时每页显示的记录数。 **取值范围**：1-50的整数，默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// **参数说明**：上一次分页查询结果中最后一条记录的ID，在上一次分页查询时由物联网平台返回获得。分页查询时物联网平台是按marker也就是记录ID降序查询的，越新的数据记录ID也会越大。若填写marker，则本次只查询记录ID小于marker的数据记录。若不填写，则从记录ID最大也就是最新的一条数据开始查询。如果需要依次查询所有数据，则每次查询时必须填写上一次查询响应中的marker值。 **取值范围**：长度为24的十六进制字符串，默认值为ffffffffffffffffffffffff。

	Marker *string `json:"marker,omitempty"`
	// **参数说明**：表示从marker后偏移offset条记录开始查询。当offset为0时，表示从marker后第一条记录开始输出。限制offset最大值是出于API性能考虑，您可以搭配marker使用该参数实现翻页，例如每页50条记录，1-11页内都可以直接使用offset跳转到指定页，但到11页后，由于offset限制为500，您需要使用第11页返回的marker作为下次查询的marker，以实现翻页到12-22页。 **取值范围**：0-500的整数，默认为0。

	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowDevicesInGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDevicesInGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowDevicesInGroupRequest", string(data)}, " ")
}
