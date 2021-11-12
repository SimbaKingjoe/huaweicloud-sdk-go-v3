package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建服务器证书请求返回对象
type CertificateInfo struct {
	// SSL证书的管理状态；暂不支持

	AdminStateUp bool `json:"admin_state_up"`
	// HTTPS协议使用的证书内容。 取值范围：PEM编码格式。

	Certificate string `json:"certificate"`
	// SSL证书的描述。

	Description string `json:"description"`
	// 服务器证书所签域名。该字段仅type为server时有效。 总长度为0-1024，由若干普通域名或泛域名组成，域名之间以\",\"分割，不超过30个域名。 普通域名：由若干字符串组成，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com； 泛域名：在普通域名的基础上仅允许首字母为\"*\"。例：*.test.com

	Domain string `json:"domain"`
	// 证书ID。

	Id string `json:"id"`
	// SSL证书的名称。

	Name string `json:"name"`
	// 服务器证书的私钥。仅type为server时有效。type为server时必选。

	PrivateKey string `json:"private_key"`
	// SSL证书的类型。分为服务器证书(server)和CA证书(client)。 默认值：server

	Type string `json:"type"`
	// 证书创建时间。

	CreatedAt string `json:"created_at"`
	// 证书更新时间。

	UpdatedAt string `json:"updated_at"`
	// 证书使用截止时间。

	ExpireTime string `json:"expire_time"`
	// 项目ID。

	ProjectId string `json:"project_id"`
}

func (o CertificateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateInfo struct{}"
	}

	return strings.Join([]string{"CertificateInfo", string(data)}, " ")
}
