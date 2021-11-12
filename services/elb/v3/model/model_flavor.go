package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 负载均衡器规格信息
type Flavor struct {
	// 规格ID。

	Id string `json:"id"`

	Info *FlavorInfo `json:"info"`
	// 规格名称。

	Name string `json:"name"`
	// 共享。

	Shared bool `json:"shared"`
	// 项目ID

	ProjectId string `json:"project_id"`
	// L4和L7 分别表示四层和七层flavor。查询支持按type过滤

	Type string `json:"type"`
	// availability_zone_ids字段，标志ELB对应L7-flavor在对应可用区是否可以售卖。 若该字段为[]代表该flavor不可售卖；若该字段为[\"ALL\"]，代表所有可用区可售卖；若仅部分可用区可售卖则返回[\"cn-north-1a\",\"cn-north-1b\"]。 可通过/v3/{project_id}/elb/availability-zones接口查询所有可售卖的可用区接口。

	AvailabilityZoneIds *[]string `json:"availability_zone_ids,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
