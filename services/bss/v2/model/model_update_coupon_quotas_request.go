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
type UpdateCouponQuotasRequest struct {
	Body *AdjustCouponQuotasReq `json:"body,omitempty"`
}

func (o UpdateCouponQuotasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateCouponQuotasRequest", string(data)}, " ")
}
