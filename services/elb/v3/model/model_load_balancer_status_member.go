package model

import (
	"encoding/json"

	"strings"
)

// lb状态树的主机组的主机状态信息
type LoadBalancerStatusMember struct {
	// provisioning的状态。 可以为：ACTIVE、PENDING_CREATE 或者ERROR。 说明：该字段为预留字段，暂未启用，默认为ACTIVE。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
	// 后端服务器ip。

	Address *string `json:"address,omitempty"`
	// 后端协议号。 取值范围[1, 65535]。

	ProtocolPort *int32 `json:"protocol_port,omitempty"`
	// 后端服务器ID。

	Id *string `json:"id,omitempty"`
	// 操作状态。 可以为：ONLINE、OFFLINE、DEGRADED、DISABLED或NO_MONITOR。默认为ONLINE。

	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusMember) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusMember struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusMember", string(data)}, " ")
}
