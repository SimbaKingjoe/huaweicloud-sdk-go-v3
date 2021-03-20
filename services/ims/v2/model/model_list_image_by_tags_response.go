package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListImageByTagsResponse struct {
	// 镜像信息列表

	Resources *[]ShowImageByTagsResource `json:"resources,omitempty"`
	// 总记录数

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListImageByTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListImageByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListImageByTagsResponse", string(data)}, " ")
}
