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
type BatchSetSubCustomerDiscountRequest struct {
	Body *SetPartnerDiscountsReq `json:"body,omitempty"`
}

func (o BatchSetSubCustomerDiscountRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchSetSubCustomerDiscountRequest", string(data)}, " ")
}
