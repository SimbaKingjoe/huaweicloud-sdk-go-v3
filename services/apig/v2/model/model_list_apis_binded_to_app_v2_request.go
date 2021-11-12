package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApisBindedToAppV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// API名称

	ApiName *string `json:"api_name,omitempty"`
	// API分组编号

	GroupId *string `json:"group_id,omitempty"`
	// API分组名称

	GroupName *string `json:"group_name,omitempty"`
	// 授权的环境编号

	EnvId *string `json:"env_id,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApisBindedToAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisBindedToAppV2Request struct{}"
	}

	return strings.Join([]string{"ListApisBindedToAppV2Request", string(data)}, " ")
}
