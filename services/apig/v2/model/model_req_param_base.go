package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ReqParamBase struct {
	// 参数名称。 长度为1 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线、英文句号组成，且只能以英文开头。

	Name string `json:"name"`
	// 参数类型

	Type ReqParamBaseType `json:"type"`
	// 参数位置

	Location ReqParamBaseLocation `json:"location"`
	// 参数默认值

	DefaultValue *string `json:"default_value,omitempty"`
	// 参数示例值

	SampleValue *string `json:"sample_value,omitempty"`
	// 是否必须 - 1：是 - 2：否  location为PATH时，required默认为1，其他场景required默认为2

	Required *int32 `json:"required,omitempty"`
	// 是否开启校验 - 1：开启校验 - 2：不开启校验

	ValidEnable *int32 `json:"valid_enable,omitempty"`
	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
	// 参数枚举值

	Enumerations *string `json:"enumerations,omitempty"`
	// 参数最小值  参数类型为NUMBER时有效

	MinNum *int32 `json:"min_num,omitempty"`
	// 参数最大值  参数类型为NUMBER时有效

	MaxNum *int32 `json:"max_num,omitempty"`
	// 参数最小长度  参数类型为STRING时有效

	MinSize *int32 `json:"min_size,omitempty"`
	// 参数最大长度  参数类型为STRING时有效

	MaxSize *int32 `json:"max_size,omitempty"`
	// 正则校验规则  暂不支持

	Regular *string `json:"regular,omitempty"`
	// JSON校验规则  暂不支持

	JsonSchema *string `json:"json_schema,omitempty"`
	// 是否透传 - 1：是 - 2：否

	PassThrough *int32 `json:"pass_through,omitempty"`
}

func (o ReqParamBase) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReqParamBase struct{}"
	}

	return strings.Join([]string{"ReqParamBase", string(data)}, " ")
}

type ReqParamBaseType struct {
	value string
}

type ReqParamBaseTypeEnum struct {
	STRING ReqParamBaseType
	NUMBER ReqParamBaseType
}

func GetReqParamBaseTypeEnum() ReqParamBaseTypeEnum {
	return ReqParamBaseTypeEnum{
		STRING: ReqParamBaseType{
			value: "STRING",
		},
		NUMBER: ReqParamBaseType{
			value: "NUMBER",
		},
	}
}

func (c ReqParamBaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ReqParamBaseType) UnmarshalJSON(b []byte) error {
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

type ReqParamBaseLocation struct {
	value string
}

type ReqParamBaseLocationEnum struct {
	PATH   ReqParamBaseLocation
	QUERY  ReqParamBaseLocation
	HEADER ReqParamBaseLocation
}

func GetReqParamBaseLocationEnum() ReqParamBaseLocationEnum {
	return ReqParamBaseLocationEnum{
		PATH: ReqParamBaseLocation{
			value: "PATH",
		},
		QUERY: ReqParamBaseLocation{
			value: "QUERY",
		},
		HEADER: ReqParamBaseLocation{
			value: "HEADER",
		},
	}
}

func (c ReqParamBaseLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ReqParamBaseLocation) UnmarshalJSON(b []byte) error {
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
