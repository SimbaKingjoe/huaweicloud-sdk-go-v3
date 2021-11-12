package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructFieldInfo struct {
	// 结构化方式

	IsAnalysis bool `json:"isAnalysis"`
	// 字段内容

	Content *string `json:"content,omitempty"`
	// 字段名称

	FieldName *string `json:"fieldName,omitempty"`
	// 字段数据类型

	Type string `json:"type"`
	// 序号

	Index *int32 `json:"index,omitempty"`
}

func (o StructFieldInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructFieldInfo struct{}"
	}

	return strings.Join([]string{"StructFieldInfo", string(data)}, " ")
}
