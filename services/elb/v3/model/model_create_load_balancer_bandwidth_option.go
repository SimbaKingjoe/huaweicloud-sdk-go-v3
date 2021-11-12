package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 带宽信息
type CreateLoadBalancerBandwidthOption struct {
	// 带宽名称

	Name string `json:"name"`
	// 带宽大小 取值范围:默认1Mbit/s~2000Mbit/s(具体范围以各区域配置为准,请参见控制台对应页面显示)。 约束:share_type是PER,该参数必须带,如果share_type是WHOLE并且id有值,该参数会忽略。 注意:调整带宽时的最小单位会根据带宽范围不同存在差异。 小于等于300Mbit/s:默认最小单位为1Mbit/s。 300Mbit/s~1000Mbit/s:默认最小单位为50Mbit/s。 大于1000Mbit/s:默认最小单位为500Mbit/s。

	Size int32 `json:"size"`
	// 按流量计费还是按带宽计费。 其中IPv6国外默认是bandwidth,国内默认是traffic。取值为traffic,表示流量计费

	ChargeMode CreateLoadBalancerBandwidthOptionChargeMode `json:"charge_mode"`
	// 有效值：PER,WHOLE 约束:其中IPv6暂不支持WHOLE类型带宽,该字段为WHOLE时,必须指定带宽ID。

	ShareType CreateLoadBalancerBandwidthOptionShareType `json:"share_type"`
	// 预留资源账单信息，默认为空表示按需计费， 非空为包周期。admin权限才能更新此字段

	BillingInfo *string `json:"billing_info,omitempty"`
}

func (o CreateLoadBalancerBandwidthOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerBandwidthOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerBandwidthOption", string(data)}, " ")
}

type CreateLoadBalancerBandwidthOptionChargeMode struct {
	value string
}

type CreateLoadBalancerBandwidthOptionChargeModeEnum struct {
	BANDWIDTH CreateLoadBalancerBandwidthOptionChargeMode
	TRAFFIC   CreateLoadBalancerBandwidthOptionChargeMode
}

func GetCreateLoadBalancerBandwidthOptionChargeModeEnum() CreateLoadBalancerBandwidthOptionChargeModeEnum {
	return CreateLoadBalancerBandwidthOptionChargeModeEnum{
		BANDWIDTH: CreateLoadBalancerBandwidthOptionChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: CreateLoadBalancerBandwidthOptionChargeMode{
			value: "traffic",
		},
	}
}

func (c CreateLoadBalancerBandwidthOptionChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerBandwidthOptionChargeMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateLoadBalancerBandwidthOptionShareType struct {
	value string
}

type CreateLoadBalancerBandwidthOptionShareTypeEnum struct {
	PER   CreateLoadBalancerBandwidthOptionShareType
	WHOLE CreateLoadBalancerBandwidthOptionShareType
}

func GetCreateLoadBalancerBandwidthOptionShareTypeEnum() CreateLoadBalancerBandwidthOptionShareTypeEnum {
	return CreateLoadBalancerBandwidthOptionShareTypeEnum{
		PER: CreateLoadBalancerBandwidthOptionShareType{
			value: "PER",
		},
		WHOLE: CreateLoadBalancerBandwidthOptionShareType{
			value: "WHOLE",
		},
	}
}

func (c CreateLoadBalancerBandwidthOptionShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerBandwidthOptionShareType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
