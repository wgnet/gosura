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
	REMOVE_REMOTE_SCHEMA_TYPE string = `remove_remote_schema`
)

type RemoveRemoteSchema struct {
	Arguments RemoveRemoteSchemaArgs `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type RemoveRemoteSchemaArgs struct {
	Name string `json:"name"`
}

type RemoveRemoteSchemaResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *RemoveRemoteSchema) SetArgs(args interface{}) error {
	switch args.(type) {
	case RemoveRemoteSchemaArgs:
		t.Arguments = args.(RemoveRemoteSchemaArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *RemoveRemoteSchema) SetVersion(version int) {
	t.Ver = version
}

func (t *RemoveRemoteSchema) SetType(name string) {
	t.QueryType = name
}

func (t *RemoveRemoteSchema) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *RemoveRemoteSchema) Method() string {
	return http.MethodPost
}

func (t *RemoveRemoteSchema) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var removeRemoteSchemaResponse RemoveRemoteSchemaResponse
	if err := json.Unmarshal(body, &removeRemoteSchemaResponse); err != nil {
		return nil, err
	}
	return removeRemoteSchemaResponse, nil
}

func NewRemoveRemoteSchemaQuery() Query {
	query := RemoveRemoteSchema{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: REMOVE_REMOTE_SCHEMA_TYPE,
	}

	return Query(&query)
}
