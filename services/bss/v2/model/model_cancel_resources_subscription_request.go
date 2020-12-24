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
type CancelResourcesSubscriptionRequest struct {
	Body *UnsubscribeResourcesReq `json:"body,omitempty"`
}

func (o CancelResourcesSubscriptionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelResourcesSubscriptionRequest", string(data)}, " ")
}
