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
type RenewalResourcesRequest struct {
	Body *RenewalResourcesReq `json:"body,omitempty"`
}

func (o RenewalResourcesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RenewalResourcesRequest", string(data)}, " ")
}
