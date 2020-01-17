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
	CREATE_QUERY_COLLECTION_TYPE string = `create_query_collection`
)

type CreateQueryCollection struct {
	Arguments CreateQueryCollectionArgs `json:"args"`
	Ver       int                       `json:"version"`
	QueryType string                    `json:"type"`
}

type CreateQueryCollectionArgs struct {
	Name       string                           `json:"name"`
	Definition QueryCollectionDefinitionQueries `json:"definition"`
	Comment    string                           `json:"comment,omitempty"`
}

type QueryCollectionDefinitionQueries struct {
	Queries []QueryCollection `json:"queries"`
}

type QueryCollection struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type CreateQueryCollectionResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *CreateQueryCollection) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateQueryCollectionArgs:
		t.Arguments = args.(CreateQueryCollectionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateQueryCollection) SetVersion(version int) {
	t.Ver = version
}

func (t *CreateQueryCollection) SetType(name string) {
	t.QueryType = name
}

func (t *CreateQueryCollection) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateQueryCollection) Method() string {
	return http.MethodPost
}

func (t *CreateQueryCollection) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createQueryCollectionResponse CreateQueryCollectionResponse
	if err := json.Unmarshal(body, &createQueryCollectionResponse); err != nil {
		return nil, err
	}
	return createQueryCollectionResponse, nil
}

func NewCreateQueryCollectionQuery() Query {
	query := CreateQueryCollection{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: CREATE_QUERY_COLLECTION_TYPE,
	}

	return Query(&query)
}
