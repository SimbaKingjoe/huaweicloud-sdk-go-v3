/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RemoveProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveProjectResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RemoveProjectResponse", string(data)}, " ")
}
