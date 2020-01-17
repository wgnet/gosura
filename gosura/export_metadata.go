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
	EXPORT_METADATA_TYPE string = `export_metadata`
)

type ExportMetadata struct {
	Arguments map[string]interface{} `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type ExportMetadataResponse map[string]interface{}

func (t *ExportMetadata) SetArgs(args interface{}) error {
	return nil
}

func (t *ExportMetadata) SetVersion(version int) {
	t.Ver = version
}

func (t *ExportMetadata) SetType(name string) {
	t.QueryType = name
}

func (t *ExportMetadata) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *ExportMetadata) Method() string {
	return http.MethodPost
}

func (t *ExportMetadata) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var exportMetadataResponse ExportMetadataResponse
	if err := json.Unmarshal(body, &exportMetadataResponse); err != nil {
		return nil, err
	}
	return exportMetadataResponse, nil
}

func NewExportMetadataQuery() Query {
	query := ExportMetadata{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: EXPORT_METADATA_TYPE,
		Arguments: make(map[string]interface{}),
	}

	return Query(&query)
}
