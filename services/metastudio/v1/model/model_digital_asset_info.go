package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DigitalAssetInfo 资产信息。
type DigitalAssetInfo struct {

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产描述。
	AssetDescription *string `json:"asset_description,omitempty"`

	// 资产创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 资产更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 资产类型。  公共资产类型： * VOICE_MODEL：音色模型 * VIDEO：视频文件 * IMAGE：图片文件 * PPT：幻灯片文件 * MUSIC: 音乐 * AUDIO: 音频 * COMMON_FILE：通用文件  分身数字人资产类型： * HUMAN_MODEL_2D：分身数字人模型 * BUSINESS_CARD_TEMPLET: 数字人名片模板  3D数字人资产类型： * HUMAN_MODEL：3D数字人模型 * SCENE：场景模型 * ANIMATION：动作动画 * MATERIAL：风格化素材 * NORMAL_MODEL: 普通模型
	AssetType *DigitalAssetInfoAssetType `json:"asset_type,omitempty"`

	// 资产状态。 * CREATING：资产创建中，主文件尚未上传 * FAILED：主文件上传失败 * UNACTIVED：主文件上传成功，资产未激活，资产不可用于其他业务（用户可更新状态） * ACTIVED：主文件上传成功，资产激活，资产可用于其他业务（用户可更新状态） * DELETING：资产删除中，资产不可用，资产可恢复 * DELETED：资产文件已删除，资产不可用，资产不可恢复 * BLOCK: 资产被冻结，资产不可用，不可查看文件。
	AssetState *DigitalAssetInfoAssetState `json:"asset_state,omitempty"`

	// 失败原因。 * AUTOMATIC_REVIEW_REJECT：自动审核失败 * MANUAL_REVIEW_REJECT：人工审核失败
	FailType *DigitalAssetInfoFailType `json:"fail_type,omitempty"`

	// 冻结/解冻/失败 原因。
	Reason *string `json:"reason,omitempty"`

	// 是否需要资产库生成封面图片。 > * 当前支持自动生成封面图片的资产类型包括VIDEO
	IsNeedGenerateCover *bool `json:"is_need_generate_cover,omitempty"`

	// 标签列表。
	Tags *[]string `json:"tags,omitempty"`

	AssetExtraMeta *AssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 设置系统属性。
	SystemProperties *[]SystemProperty `json:"system_properties,omitempty"`

	// 资产下的文件。
	Files *[]AssetFileInfo `json:"files,omitempty"`
}

func (o DigitalAssetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DigitalAssetInfo struct{}"
	}

	return strings.Join([]string{"DigitalAssetInfo", string(data)}, " ")
}

type DigitalAssetInfoAssetType struct {
	value string
}

type DigitalAssetInfoAssetTypeEnum struct {
	HUMAN_MODEL           DigitalAssetInfoAssetType
	VOICE_MODEL           DigitalAssetInfoAssetType
	SCENE                 DigitalAssetInfoAssetType
	ANIMATION             DigitalAssetInfoAssetType
	VIDEO                 DigitalAssetInfoAssetType
	IMAGE                 DigitalAssetInfoAssetType
	PPT                   DigitalAssetInfoAssetType
	MATERIAL              DigitalAssetInfoAssetType
	NORMAL_MODEL          DigitalAssetInfoAssetType
	COMMON_FILE           DigitalAssetInfoAssetType
	HUMAN_MODEL_2_D       DigitalAssetInfoAssetType
	BUSINESS_CARD_TEMPLET DigitalAssetInfoAssetType
	MUSIC                 DigitalAssetInfoAssetType
	AUDIO                 DigitalAssetInfoAssetType
}

func GetDigitalAssetInfoAssetTypeEnum() DigitalAssetInfoAssetTypeEnum {
	return DigitalAssetInfoAssetTypeEnum{
		HUMAN_MODEL: DigitalAssetInfoAssetType{
			value: "HUMAN_MODEL",
		},
		VOICE_MODEL: DigitalAssetInfoAssetType{
			value: "VOICE_MODEL",
		},
		SCENE: DigitalAssetInfoAssetType{
			value: "SCENE",
		},
		ANIMATION: DigitalAssetInfoAssetType{
			value: "ANIMATION",
		},
		VIDEO: DigitalAssetInfoAssetType{
			value: "VIDEO",
		},
		IMAGE: DigitalAssetInfoAssetType{
			value: "IMAGE",
		},
		PPT: DigitalAssetInfoAssetType{
			value: "PPT",
		},
		MATERIAL: DigitalAssetInfoAssetType{
			value: "MATERIAL",
		},
		NORMAL_MODEL: DigitalAssetInfoAssetType{
			value: "NORMAL_MODEL",
		},
		COMMON_FILE: DigitalAssetInfoAssetType{
			value: "COMMON_FILE",
		},
		HUMAN_MODEL_2_D: DigitalAssetInfoAssetType{
			value: "HUMAN_MODEL_2D",
		},
		BUSINESS_CARD_TEMPLET: DigitalAssetInfoAssetType{
			value: "BUSINESS_CARD_TEMPLET",
		},
		MUSIC: DigitalAssetInfoAssetType{
			value: "MUSIC",
		},
		AUDIO: DigitalAssetInfoAssetType{
			value: "AUDIO",
		},
	}
}

func (c DigitalAssetInfoAssetType) Value() string {
	return c.value
}

func (c DigitalAssetInfoAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalAssetInfoAssetType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type DigitalAssetInfoAssetState struct {
	value string
}

type DigitalAssetInfoAssetStateEnum struct {
	CREATING  DigitalAssetInfoAssetState
	FAILED    DigitalAssetInfoAssetState
	UNACTIVED DigitalAssetInfoAssetState
	ACTIVED   DigitalAssetInfoAssetState
	DELETING  DigitalAssetInfoAssetState
	DELETED   DigitalAssetInfoAssetState
	BLOCK     DigitalAssetInfoAssetState
}

func GetDigitalAssetInfoAssetStateEnum() DigitalAssetInfoAssetStateEnum {
	return DigitalAssetInfoAssetStateEnum{
		CREATING: DigitalAssetInfoAssetState{
			value: "CREATING",
		},
		FAILED: DigitalAssetInfoAssetState{
			value: "FAILED",
		},
		UNACTIVED: DigitalAssetInfoAssetState{
			value: "UNACTIVED",
		},
		ACTIVED: DigitalAssetInfoAssetState{
			value: "ACTIVED",
		},
		DELETING: DigitalAssetInfoAssetState{
			value: "DELETING",
		},
		DELETED: DigitalAssetInfoAssetState{
			value: "DELETED",
		},
		BLOCK: DigitalAssetInfoAssetState{
			value: "BLOCK",
		},
	}
}

func (c DigitalAssetInfoAssetState) Value() string {
	return c.value
}

func (c DigitalAssetInfoAssetState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalAssetInfoAssetState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type DigitalAssetInfoFailType struct {
	value string
}

type DigitalAssetInfoFailTypeEnum struct {
	AUTOMATIC_REVIEW_REJECT DigitalAssetInfoFailType
	MANUAL_REVIEW_REJECT    DigitalAssetInfoFailType
}

func GetDigitalAssetInfoFailTypeEnum() DigitalAssetInfoFailTypeEnum {
	return DigitalAssetInfoFailTypeEnum{
		AUTOMATIC_REVIEW_REJECT: DigitalAssetInfoFailType{
			value: "AUTOMATIC_REVIEW_REJECT",
		},
		MANUAL_REVIEW_REJECT: DigitalAssetInfoFailType{
			value: "MANUAL_REVIEW_REJECT",
		},
	}
}

func (c DigitalAssetInfoFailType) Value() string {
	return c.value
}

func (c DigitalAssetInfoFailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DigitalAssetInfoFailType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
