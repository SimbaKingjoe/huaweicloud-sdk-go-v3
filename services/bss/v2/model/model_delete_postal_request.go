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
type DeletePostalRequest struct {
	AddressId string `json:"address_id"`
}

func (o DeletePostalRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePostalRequest", string(data)}, " ")
}
