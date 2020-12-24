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
type CheckUserIdentityRequest struct {
	Body *CheckSubcustomerUserReq `json:"body,omitempty"`
}

func (o CheckUserIdentityRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CheckUserIdentityRequest", string(data)}, " ")
}
