package model

import (
	"encoding/json"

	"strings"
)

type KeystoneListUsersResult struct {
	// IAM用户密码状态。true：需要修改密码，false：正常。

	PwdStatus *bool `json:"pwd_status,omitempty"`
	// IAM用户所属账号ID。

	DomainId string `json:"domain_id"`
	// IAM用户退出系统前，在控制台最后访问的项目ID。

	LastProjectId *string `json:"last_project_id,omitempty"`
	// IAM用户名。

	Name string `json:"name"`
	// IAM用户描述信息。

	Description *string `json:"description,omitempty"`
	// IAM用户密码过期时间（UTC时间），“null”表示密码不过期。

	PasswordExpiresAt string `json:"password_expires_at"`

	Links *Links `json:"links"`
	// IAM用户ID。

	Id string `json:"id"`
	// IAM用户是否启用。true表示启用，false表示停用，默认为true。

	Enabled bool `json:"enabled"`
	// IAM用户的密码强度。high：密码强度高；mid：密码强度中等；low：密码强度低。

	PwdStrength *string `json:"pwd_strength,omitempty"`
}

func (o KeystoneListUsersResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListUsersResult struct{}"
	}

	return strings.Join([]string{"KeystoneListUsersResult", string(data)}, " ")
}
