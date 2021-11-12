package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateRequestThrottlingPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *ThrottleBindingReq `json:"body,omitempty"`
}

func (o AssociateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"AssociateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
