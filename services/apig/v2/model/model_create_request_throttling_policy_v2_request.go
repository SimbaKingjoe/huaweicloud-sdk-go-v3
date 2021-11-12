package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRequestThrottlingPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *ThrottleReq `json:"body,omitempty"`
}

func (o CreateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"CreateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
