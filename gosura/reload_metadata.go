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
	RELOAD_METADATA_TYPE string = `reload_metadata`
)

type ReloadMetadata struct {
	Arguments map[string]interface{} `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type ReloadMetadataResponse map[string]interface{}

// SetArgs do nothing here
func (t *ReloadMetadata) SetArgs(args interface{}) error {
	return nil
}

func (t *ReloadMetadata) SetVersion(version int) {
	t.Ver = version
}

func (t *ReloadMetadata) SetType(name string) {
	t.QueryType = name
}

func (t *ReloadMetadata) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *ReloadMetadata) Method() string {
	return http.MethodPost
}

func (t *ReloadMetadata) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var reloadMetadataResponse ReloadMetadataResponse
	if err := json.Unmarshal(body, &reloadMetadataResponse); err != nil {
		return nil, err
	}
	return reloadMetadataResponse, nil
}

func NewReloadMetadataQuery() Query {
	query := ReloadMetadata{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: RELOAD_METADATA_TYPE,
		Arguments: make(map[string]interface{}),
	}

	return Query(&query)
}
