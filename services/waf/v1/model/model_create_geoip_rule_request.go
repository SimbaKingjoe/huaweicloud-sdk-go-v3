package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGeoipRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`

	Body *CreateGeoIpRuleRequestBody `json:"body,omitempty"`
}

func (o CreateGeoipRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGeoipRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateGeoipRuleRequest", string(data)}, " ")
}
