package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDialogUrlRequest Request Object
type CreateDialogUrlRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *CreateDialogUrlReq `json:"body,omitempty"`
}

func (o CreateDialogUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDialogUrlRequest struct{}"
	}

	return strings.Join([]string{"CreateDialogUrlRequest", string(data)}, " ")
}
