package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateProviderVersionRequest Request Object
type CreatePrivateProviderVersionRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 私有provider（private-provider）的名称。此名字在domain_id+region下应唯一，可以使用小写英文、数字、中划线。仅支持以小写英文、数字开头结尾。
	ProviderName string `json:"provider_name"`

	Body *CreatePrivateProviderVersionRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateProviderVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateProviderVersionRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateProviderVersionRequest", string(data)}, " ")
}
