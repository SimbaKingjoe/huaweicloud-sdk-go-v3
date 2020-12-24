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
type ListCitiesRequest struct {
	XLanguage    *string `json:"X-Language,omitempty"`
	ProvinceCode string  `json:"province_code"`
	Offset       *int32  `json:"offset,omitempty"`
	Limit        *int32  `json:"limit,omitempty"`
}

func (o ListCitiesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCitiesRequest", string(data)}, " ")
}
