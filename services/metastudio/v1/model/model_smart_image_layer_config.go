package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmartImageLayerConfig 素材图片图层配置。
type SmartImageLayerConfig struct {

	// 图片文件的URL。
	ImageUrl string `json:"image_url"`

	// 图片显示时长。单位s * 0表示一直显示。
	DisplayDuration *int32 `json:"display_duration,omitempty"`
}

func (o SmartImageLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartImageLayerConfig struct{}"
	}

	return strings.Join([]string{"SmartImageLayerConfig", string(data)}, " ")
}
