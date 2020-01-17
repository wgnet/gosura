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
	DROP_QUERY_FROM_COLLECTION_TYPE string = `drop_query_from_collection`
)

type DropQueryFromCollection struct {
	Arguments DropQueryFromCollectionArgs `json:"args"`
	Ver       int                         `json:"version"`
	QueryType string                      `json:"type"`
}

type DropQueryFromCollectionArgs struct {
	Collection string `json:"collection_name"`
	Name       string `json:"query_name"`
}

type DropQueryFromCollectionResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *DropQueryFromCollection) SetArgs(args interface{}) error {
	switch args.(type) {
	case DropQueryFromCollectionArgs:
		t.Arguments = args.(DropQueryFromCollectionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *DropQueryFromCollection) SetVersion(version int) {
	t.Ver = version
}

func (t *DropQueryFromCollection) SetType(name string) {
	t.QueryType = name
}

func (t *DropQueryFromCollection) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *DropQueryFromCollection) Method() string {
	return http.MethodPost
}

func (t *DropQueryFromCollection) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var dropQueryFromCollectionResponse DropQueryFromCollectionResponse
	if err := json.Unmarshal(body, &dropQueryFromCollectionResponse); err != nil {
		return nil, err
	}
	return dropQueryFromCollectionResponse, nil
}

func NewDropQueryFromCollectionQuery() Query {
	query := DropQueryFromCollection{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: DROP_QUERY_FROM_COLLECTION_TYPE,
	}

	return Query(&query)
}
