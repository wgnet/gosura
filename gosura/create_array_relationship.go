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
	CREATE_ARRAY_RELATIONSHIP_TYPE string = `create_array_relationship`
)

type CreateArrayRelationship struct {
	Arguments CreateArrayRelationshipArgs `json:"args"`
	Ver       int                         `json:"version"`
	QueryType string                      `json:"type"`
}

type CreateArrayRelationshipArgs struct {
	Table string      `json:"table"`
	Name  string      `json:"name"`
	Using interface{} `json:"using"`
}

type CreateArrayRelationshipUsingAuto struct {
	ForeignKeyConstraintOn CreateArrayRelationshipUsingAutoConstraint `json:"foreign_key_constraint_on"`
}

type CreateArrayRelationshipUsingAutoConstraint struct {
	Table  string `json:"table"`
	Column string `json:"column"`
}

type CreateArrayRelationshipResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *CreateArrayRelationship) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateArrayRelationshipArgs:
		t.Arguments = args.(CreateArrayRelationshipArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateArrayRelationship) SetVersion(version int) {
	t.Ver = version
}

func (t *CreateArrayRelationship) SetType(name string) {
	t.QueryType = name
}

func (t *CreateArrayRelationship) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateArrayRelationship) Method() string {
	return http.MethodPost
}

func (t *CreateArrayRelationship) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createArrayRelationshipResponse CreateArrayRelationshipResponse
	if err := json.Unmarshal(body, &createArrayRelationshipResponse); err != nil {
		return nil, err
	}
	return createArrayRelationshipResponse, nil
}

func NewCreateArrayRelationshipQuery() Query {
	query := CreateArrayRelationship{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: CREATE_ARRAY_RELATIONSHIP_TYPE,
	}

	return Query(&query)
}
