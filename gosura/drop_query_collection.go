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
	DROP_QUERY_COLLECTION_TYPE string = `drop_query_collection`
)

type DropQueryCollection struct {
	Arguments DropQueryCollectionArgs `json:"args"`
	Ver       int                     `json:"version"`
	QueryType string                  `json:"type"`
}

type DropQueryCollectionArgs struct {
	Collection string `json:"collection"`
	Cascade    bool   `json:"cascade"`
}

type DropQueryCollectionResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *DropQueryCollection) SetArgs(args interface{}) error {
	switch args.(type) {
	case DropQueryCollectionArgs:
		t.Arguments = args.(DropQueryCollectionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *DropQueryCollection) SetVersion(version int) {
	t.Ver = version
}

func (t *DropQueryCollection) SetType(name string) {
	t.QueryType = name
}

func (t *DropQueryCollection) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *DropQueryCollection) Method() string {
	return http.MethodPost
}

func (t *DropQueryCollection) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var dropQueryCollectionResponse DropQueryCollectionResponse
	if err := json.Unmarshal(body, &dropQueryCollectionResponse); err != nil {
		return nil, err
	}
	return dropQueryCollectionResponse, nil
}

func NewDropQueryCollectionQuery() Query {
	query := DropQueryCollection{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: DROP_QUERY_COLLECTION_TYPE,
	}

	return Query(&query)
}
