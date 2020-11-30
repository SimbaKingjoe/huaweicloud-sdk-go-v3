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
type DeleteVpcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVpcResponse", string(data)}, " ")
}
