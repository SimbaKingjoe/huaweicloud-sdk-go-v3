package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleBindingThrottleResp struct {
	// 流控策略的ID

	Id *string `json:"id,omitempty"`
	// 流控策略的名称

	Name *string `json:"name,omitempty"`
	// 单个API流控时间内能够被访问的次数限制

	ApiCallLimits *int32 `json:"api_call_limits,omitempty"`
	// 单个用户流控时间内能够访问API的次数限制

	UserCallLimits *int32 `json:"user_call_limits,omitempty"`
	// 单个APP流控时间内能够访问API的次数限制

	AppCallLimits *int32 `json:"app_call_limits,omitempty"`
	// 单个源IP流控时间内能够访问API的次数限制

	IpCallLimits *int32 `json:"ip_call_limits,omitempty"`
	// 流控的时长

	TimeInterval *int32 `json:"time_interval,omitempty"`
	// 流控的时间单位

	TimeUnit *string `json:"time_unit,omitempty"`
	// 描述

	Remark *string `json:"remark,omitempty"`
	// 创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 是否包含特殊流控 - 1：包含 - 2：不包含

	IsIncludeSpecialThrottle *int32 `json:"is_include_special_throttle,omitempty"`
	// 流控策略生效的环境（即在哪个环境上有效）

	EnvName *string `json:"env_name,omitempty"`
	// 流控策略的类型

	Type *int32 `json:"type,omitempty"`
	// 流控策略与API绑定关系编号

	BindId *string `json:"bind_id,omitempty"`
	// 流控策略与API绑定时间

	BindTime *sdktime.SdkTime `json:"bind_time,omitempty"`
	// 流控策略绑定的API数量

	BindNum *int32 `json:"bind_num,omitempty"`
	// 是否开启动态流控，暂不支持

	EnableAdaptiveControl *string `json:"enable_adaptive_control,omitempty"`
}

func (o ThrottleBindingThrottleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleBindingThrottleResp struct{}"
	}

	return strings.Join([]string{"ThrottleBindingThrottleResp", string(data)}, " ")
}
