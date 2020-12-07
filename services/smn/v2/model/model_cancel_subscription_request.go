/*
 * smn
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelSubscriptionRequest struct {
	SubscriptionUrn string `json:"subscription_urn"`
}

func (o CancelSubscriptionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelSubscriptionRequest", string(data)}, " ")
}
