package model

import (
	"encoding/json"

	"strings"
)

type VaultAssociate struct {
	// 目标vault ID , 只有设置复制策略时使用，而且必传
	DestinationVaultId *string `json:"destination_vault_id,omitempty"`
	// 策略ID
	PolicyId string `json:"policy_id"`
}

func (o VaultAssociate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultAssociate struct{}"
	}

	return strings.Join([]string{"VaultAssociate", string(data)}, " ")
}
