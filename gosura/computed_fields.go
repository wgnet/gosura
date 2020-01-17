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
	ADD_COMPUTED_FIELD_TYPE  string = `add_computed_field`
	DROP_COMPUTED_FIELD_TYPE string = `drop_computed_field`
)

type ComputedField struct {
	Arguments ComputedFieldArgs `json:"args"`
	Ver       int               `json:"version"`
	QueryType string            `json:"type"`
}

type ComputedFieldArgs struct {
	Table      TableArgs                   `json:"table"`
	Name       string                      `json:"name"`
	TableArg   string                      `json:"table_argument,omitempty"`
	Definition ComputedFieldArgsDefinition `json:"definition,omitempty"`
	Comment    string                      `json:"comment,omitempty"`
	Cascade    bool                        `json:"cascade,omitempty"`
}

type ComputedFieldArgsDefinition struct {
	Function TableArgs `json:"function"`
}

type ComputedFieldResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *ComputedField) SetArgs(args interface{}) error {
	switch args.(type) {
	case ComputedFieldArgs:
		t.Arguments = args.(ComputedFieldArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *ComputedField) SetVersion(version int) {
	t.Ver = version
}

func (t *ComputedField) SetType(name string) {
	t.QueryType = name
}

func (t *ComputedField) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *ComputedField) Method() string {
	return http.MethodPost
}

func (t *ComputedField) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var computedFieldResponse ComputedFieldResponse
	if err := json.Unmarshal(body, &computedFieldResponse); err != nil {
		return nil, err
	}
	return computedFieldResponse, nil
}

func NewComputedFieldQuery(drop bool) Query {
	queryType := ADD_COMPUTED_FIELD_TYPE
	if drop {
		queryType = DROP_COMPUTED_FIELD_TYPE
	}
	query := ComputedField{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: queryType,
	}

	return Query(&query)
}
