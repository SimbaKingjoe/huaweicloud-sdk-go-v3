/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSubEnterpriseAccountRequest struct {
	Body *CreateSubCustomerReqV2 `json:"body,omitempty"`
}

func (o CreateSubEnterpriseAccountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSubEnterpriseAccountRequest", string(data)}, " ")
}
