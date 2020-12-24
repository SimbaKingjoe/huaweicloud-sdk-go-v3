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

// Response Object
type ReclaimPartnerCouponsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ReclaimPartnerCouponsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ReclaimPartnerCouponsResponse", string(data)}, " ")
}
