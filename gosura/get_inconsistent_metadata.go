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
	"net/http"
)

const (
	GET_INCONSISTENT_METADATA_TYPE string = `get_inconsistent_metadata`
)

type GetInconsistentMetadata struct {
	Arguments map[string]interface{} `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type GetInconsistentMetadataResponse []map[string]interface{}

// SetArgs do nothing here
func (t *GetInconsistentMetadata) SetArgs(args interface{}) error {
	return nil
}

func (t *GetInconsistentMetadata) SetVersion(version int) {
	t.Ver = version
}

func (t *GetInconsistentMetadata) SetType(name string) {
	t.QueryType = name
}

func (t *GetInconsistentMetadata) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *GetInconsistentMetadata) Method() string {
	return http.MethodPost
}

func (t *GetInconsistentMetadata) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var getInconsistentMetadataResponse GetInconsistentMetadataResponse
	if err := json.Unmarshal(body, &getInconsistentMetadataResponse); err != nil {
		return nil, err
	}
	return getInconsistentMetadataResponse, nil
}

func NewGetInconsistentMetadataQuery() Query {
	query := GetInconsistentMetadata{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: GET_INCONSISTENT_METADATA_TYPE,
		Arguments: make(map[string]interface{}),
	}

	return Query(&query)
}
