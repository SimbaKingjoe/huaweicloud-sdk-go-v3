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
type ListSubCustomerDiscountsRequest struct {
	CustomerId        string  `json:"customer_id"`
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListSubCustomerDiscountsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSubCustomerDiscountsRequest", string(data)}, " ")
}
