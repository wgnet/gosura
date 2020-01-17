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
	CREATE_INSERT_PERMISSION_TYPE string = `create_insert_permission`
)

type CreateInsertPermission struct {
	Arguments CreateInsertPermissionArgs `json:"args"`
	QueryType string                     `json:"type"`
}

type CreateInsertPermissionArgs struct {
	Table      string            `json:"table"`
	Role       string            `json:"role"`
	Permission *InsertPermission `json:"permission"`
	Comment    string            `json:"comment,omitempty"`
}

type CreateInsertPermissionResponse map[string]interface{}

func (t *CreateInsertPermission) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateInsertPermissionArgs:
		t.Arguments = args.(CreateInsertPermissionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateInsertPermission) SetVersion(_ int) {}

func (t *CreateInsertPermission) SetType(name string) {
	t.QueryType = name
}

func (t *CreateInsertPermission) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateInsertPermission) Method() string {
	return http.MethodPost
}

func (t *CreateInsertPermission) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createInsertPermissionResponse CreateInsertPermissionResponse
	if err := json.Unmarshal(body, &createInsertPermissionResponse); err != nil {
		return nil, err
	}
	return createInsertPermissionResponse, nil
}

func NewCreateInsertPermissionQuery() Query {
	query := CreateInsertPermission{
		QueryType: CREATE_INSERT_PERMISSION_TYPE,
	}

	return Query(&query)
}
