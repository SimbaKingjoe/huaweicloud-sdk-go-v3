package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleBindingReq struct {
	// 流控策略编号

	StrategyId string `json:"strategy_id"`
	// API的发布记录编号

	PublishIds []string `json:"publish_ids"`
}

func (o ThrottleBindingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleBindingReq struct{}"
	}

	return strings.Join([]string{"ThrottleBindingReq", string(data)}, " ")
}
