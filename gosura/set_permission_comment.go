/**
 * Copyright 2019-2020 Wargaming Group Limited
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
**/

package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SET_PERMISSION_COMMENT_TYPE string = `set_permission_comment`
)

type SetPermissionComment struct {
	Arguments SetPermissionCommentArgs `json:"args"`
	QueryType string                   `json:"type"`
}

type SetPermissionCommentArgs struct {
	Table   string `json:"table"`
	Role    string `json:"role"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
}

type SetPermissionCommentResponse map[string]interface{}

func (t *SetPermissionComment) SetArgs(args interface{}) error {
	switch args.(type) {
	case SetPermissionCommentArgs:
		t.Arguments = args.(SetPermissionCommentArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *SetPermissionComment) SetVersion(_ int) {}

func (t *SetPermissionComment) SetType(name string) {
	t.QueryType = name
}

func (t *SetPermissionComment) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *SetPermissionComment) Method() string {
	return http.MethodPost
}

func (t *SetPermissionComment) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var setPermissionCommentResponse SetPermissionCommentResponse
	if err := json.Unmarshal(body, &setPermissionCommentResponse); err != nil {
		return nil, err
	}
	return setPermissionCommentResponse, nil
}

func NewSetPermissionCommentQuery() Query {
	query := SetPermissionComment{
		QueryType: SET_PERMISSION_COMMENT_TYPE,
	}

	return Query(&query)
}
