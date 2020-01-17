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
	CREATE_UPDATE_PERMISSION_TYPE string = `create_update_permission`
)

type CreateUpdatePermission struct {
	Arguments CreateUpdatePermissionArgs `json:"args"`
	QueryType string                     `json:"type"`
}

type CreateUpdatePermissionArgs struct {
	Table      string            `json:"table"`
	Role       string            `json:"role"`
	Permission *UpdatePermission `json:"permission"`
	Comment    string            `json:"comment,omitempty"`
}

type CreateUpdatePermissionResponse map[string]interface{}

func (t *CreateUpdatePermission) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateUpdatePermissionArgs:
		t.Arguments = args.(CreateUpdatePermissionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateUpdatePermission) SetVersion(_ int) {}

func (t *CreateUpdatePermission) SetType(name string) {
	t.QueryType = name
}

func (t *CreateUpdatePermission) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateUpdatePermission) Method() string {
	return http.MethodPost
}

func (t *CreateUpdatePermission) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createUpdatePermissionResponse CreateUpdatePermissionResponse
	if err := json.Unmarshal(body, &createUpdatePermissionResponse); err != nil {
		return nil, err
	}
	return createUpdatePermissionResponse, nil
}

func NewCreateUpdatePermissionQuery() Query {
	query := CreateUpdatePermission{
		QueryType: CREATE_UPDATE_PERMISSION_TYPE,
	}

	return Query(&query)
}
