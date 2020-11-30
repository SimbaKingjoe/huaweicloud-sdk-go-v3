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
type NeutronDeleteFirewallPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFirewallPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronDeleteFirewallPolicyResponse", string(data)}, " ")
}
