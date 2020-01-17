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
	DROP_INCONSISTENT_METADATA_TYPE string = `drop_inconsistent_metadata`
)

type DropInconsistentMetadata struct {
	Arguments map[string]interface{} `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type DropInconsistentMetadataResponse map[string]interface{}

// SetArgs do nothing here
func (t *DropInconsistentMetadata) SetArgs(args interface{}) error {
	return nil
}

func (t *DropInconsistentMetadata) SetVersion(version int) {
	t.Ver = version
}

func (t *DropInconsistentMetadata) SetType(name string) {
	t.QueryType = name
}

func (t *DropInconsistentMetadata) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *DropInconsistentMetadata) Method() string {
	return http.MethodPost
}

func (t *DropInconsistentMetadata) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var dropInconsistentMetadataResponse DropInconsistentMetadataResponse
	if err := json.Unmarshal(body, &dropInconsistentMetadataResponse); err != nil {
		return nil, err
	}
	return dropInconsistentMetadataResponse, nil
}

func NewDropInconsistentMetadataQuery() Query {
	query := DropInconsistentMetadata{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: DROP_INCONSISTENT_METADATA_TYPE,
		Arguments: make(map[string]interface{}),
	}

	return Query(&query)
}
