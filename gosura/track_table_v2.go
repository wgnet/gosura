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

type TrackTableV2 struct {
	Arguments TrackTableV2Args `json:"args"`
	Ver       int              `json:"version"`
	QueryType string           `json:"type"`
}

type TrackTableV2Args struct {
	Table  string             `json:"table"`
	Config TrackTableV2Config `json:"configuration"`
}

type TrackTableV2Config struct {
	CustomRootFields  CustomRootFields  `json:"custom_root_fields"`
	CustomColumnNames map[string]string `json:"custom_column_names"`
}

type CustomRootFields struct {
	Select          string `json:"select"`
	SelectByPk      string `json:"select_by_pk"`
	SelectAggregate string `json:"select_aggregate"`
	Insert          string `json:"insert"`
	Update          string `json:"update"`
	Delete          string `json:"delete"`
}

type TrackTableV2Response struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *TrackTableV2) SetArgs(args interface{}) error {
	switch args.(type) {
	case TrackTableV2Args:
		t.Arguments = args.(TrackTableV2Args)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *TrackTableV2) SetVersion(version int) {
	t.Ver = version
}

func (t *TrackTableV2) SetType(name string) {
	t.QueryType = name
}

func (t *TrackTableV2) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TrackTableV2) Method() string {
	return http.MethodPost
}

func (t *TrackTableV2) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var trackTableResponse TrackTableV2Response
	if err := json.Unmarshal(body, &trackTableResponse); err != nil {
		return nil, err
	}
	return trackTableResponse, nil
}

func NewTrackTableV2Query() Query {
	query := TrackTableV2{
		Ver:       QUERY_VERSION_V2,
		QueryType: TRACK_TABLE_TYPE,
	}

	return Query(&query)
}
