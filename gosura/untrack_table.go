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
	UNTRACK_TABLE_TYPE string = `untrack_table`
)

type UntrackTable struct {
	Arguments UntrackTableArgs `json:"args"`
	Ver       int              `json:"version"`
	QueryType string           `json:"type"`
}

type UntrackTableArgs struct {
	Table   TableArgs `json:"table"`
	Cascade bool      `json:"cascade"`
}

type UntrackTableResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *UntrackTable) SetArgs(args interface{}) error {
	switch args.(type) {
	case UntrackTableArgs:
		t.Arguments = args.(UntrackTableArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *UntrackTable) SetVersion(version int) {
	t.Ver = version
}

func (t *UntrackTable) SetType(name string) {
	t.QueryType = name
}

func (t *UntrackTable) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *UntrackTable) Method() string {
	return http.MethodPost
}

func (t *UntrackTable) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var untrackTableResponse UntrackTableResponse
	if err := json.Unmarshal(body, &untrackTableResponse); err != nil {
		return nil, err
	}

	return untrackTableResponse, nil
}

func NewUntrackTableQuery() Query {
	query := UntrackTable{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: UNTRACK_TABLE_TYPE,
	}

	return Query(&query)
}
