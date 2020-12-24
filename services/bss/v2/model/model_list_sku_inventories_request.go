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
type ListSkuInventoriesRequest struct {
	Body *QuerySkuInventoriesReq `json:"body,omitempty"`
}

func (o ListSkuInventoriesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSkuInventoriesRequest", string(data)}, " ")
}
