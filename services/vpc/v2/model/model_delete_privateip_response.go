/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePrivateipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePrivateipResponse", string(data)}, " ")
}
