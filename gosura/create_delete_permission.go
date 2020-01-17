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
	CREATE_DELETE_PERMISSION_TYPE string = `create_delete_permission`
)

type CreateDeletePermission struct {
	Arguments CreateDeletePermissionArgs `json:"args"`
	QueryType string                     `json:"type"`
}

type CreateDeletePermissionArgs struct {
	Table      string            `json:"table"`
	Role       string            `json:"role"`
	Permission *DeletePermission `json:"permission"`
	Comment    string            `json:"comment,omitempty"`
}

type CreateDeletePermissionResponse map[string]interface{}

func (t *CreateDeletePermission) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateDeletePermissionArgs:
		t.Arguments = args.(CreateDeletePermissionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateDeletePermission) SetVersion(_ int) {}

func (t *CreateDeletePermission) SetType(name string) {
	t.QueryType = name
}

func (t *CreateDeletePermission) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateDeletePermission) Method() string {
	return http.MethodPost
}

func (t *CreateDeletePermission) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createDeletePermissionResponse CreateDeletePermissionResponse
	if err := json.Unmarshal(body, &createDeletePermissionResponse); err != nil {
		return nil, err
	}
	return createDeletePermissionResponse, nil
}

func NewCreateDeletePermissionQuery() Query {
	query := CreateDeletePermission{
		QueryType: CREATE_DELETE_PERMISSION_TYPE,
	}

	return Query(&query)
}
