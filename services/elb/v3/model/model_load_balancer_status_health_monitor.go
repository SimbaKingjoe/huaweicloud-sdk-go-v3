package model

import (
	"encoding/json"

	"strings"
)

// lb状态树的主机组健康检查器状态信息
type LoadBalancerStatusHealthMonitor struct {
	// 类型，可以为TCP、UDP_CONNECT或HTTP。

	Type *string `json:"type,omitempty"`
	// 健康检查ID。

	Id *string `json:"id,omitempty"`
	// 健康检查名称。

	Name *string `json:"name,omitempty"`
	// provisioning的状态。 可以为：ACTIVE、PENDING_CREATE 或者ERROR。默认为ACTIVE。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
}

func (o LoadBalancerStatusHealthMonitor) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusHealthMonitor struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusHealthMonitor", string(data)}, " ")
}
