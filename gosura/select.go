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
	SELECT_QUERY_TYPE string = `select`
)

type Select struct {
	QueryType string     `json:"type"`
	Args      SelectArgs `json:"args"`
}

type SelectArgs struct {
	Table   TableArgs       `json:"table"`
	Columns []*SelectColumn `json:"columns"`
	OrderBy *OrderBy        `json:"order_by,omitempty"`
	Limit   int64           `json:"limit,omitempty"`
	Offset  int64           `json:"offset,omitempty"`
	Where   Bool            `json:"where,omitempty"`
}

type OrderBy struct {
	Column PGColumn `json:"column"`
	Order  string   `json:"order"`
	Nulls  string   `json:"nulls"`
}

type SelectResponse []map[string]interface{}

func (s *Select) SetType(t string) {
	s.QueryType = t
}

func (s *Select) SetVersion(v int) {}

func (s *Select) SetArgs(args interface{}) error {
	switch args.(type) {
	case SelectArgs:
		s.Args = args.(SelectArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (s *Select) Byte() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Select) Method() string {
	return http.MethodPost
}

func (t *Select) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var selectResponse SelectResponse
	if err := json.Unmarshal(body, &selectResponse); err != nil {
		return nil, err
	}
	return selectResponse, nil
}

func NewSelectQuery() Query {
	return Query(&Select{
		QueryType: SELECT_QUERY_TYPE,
	})
}
