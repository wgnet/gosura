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
	DROP_COLLECTION_FROM_ALLOWLIST_TYPE string = `drop_collection_from_allowlist_type`
)

type DropCollectionFromAllowlist struct {
	Arguments DropCollectionFromAllowlistArgs `json:"args"`
	Ver       int                             `json:"version"`
	QueryType string                          `json:"type"`
}

type DropCollectionFromAllowlistArgs struct {
	Collection string `json:"collection"`
}

type DropCollectionFromAllowlistResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *DropCollectionFromAllowlist) SetArgs(args interface{}) error {
	switch args.(type) {
	case DropCollectionFromAllowlistArgs:
		t.Arguments = args.(DropCollectionFromAllowlistArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *DropCollectionFromAllowlist) SetVersion(version int) {
	t.Ver = version
}

func (t *DropCollectionFromAllowlist) SetType(name string) {
	t.QueryType = name
}

func (t *DropCollectionFromAllowlist) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *DropCollectionFromAllowlist) Method() string {
	return http.MethodPost
}

func (t *DropCollectionFromAllowlist) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var dropCollectionFromAllowlistResponse DropCollectionFromAllowlistResponse
	if err := json.Unmarshal(body, &dropCollectionFromAllowlistResponse); err != nil {
		return nil, err
	}
	return dropCollectionFromAllowlistResponse, nil
}

func NewDropCollectionFromAllowlistQuery() Query {
	query := DropCollectionFromAllowlist{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: DROP_COLLECTION_FROM_ALLOWLIST_TYPE,
	}

	return Query(&query)
}
