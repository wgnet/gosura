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
	TRACK_FUNCTION_TYPE   string = `track_function`
	UNTRACK_FUNCTION_TYPE string = `untrack_function`
)

type TrackUntrackFunction struct {
	Arguments TrackUntrackFunctionArgs `json:"args"`
	Ver       int                      `json:"version"`
	QueryType string                   `json:"type"`
}

type TrackUntrackFunctionArgs struct {
	Schema string `json:"schema"`
	Name   string `json:"name"`
}

type TrackUntrackFunctionResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *TrackUntrackFunction) SetArgs(args interface{}) error {
	switch args.(type) {
	case TrackUntrackFunctionArgs:
		t.Arguments = args.(TrackUntrackFunctionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *TrackUntrackFunction) SetVersion(version int) {
	t.Ver = version
}

func (t *TrackUntrackFunction) SetType(name string) {
	t.QueryType = name
}

func (t *TrackUntrackFunction) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TrackUntrackFunction) Method() string {
	return http.MethodPost
}

func (t *TrackUntrackFunction) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var trackUntrackFunctionResponse TrackUntrackFunctionResponse
	if err := json.Unmarshal(body, &trackUntrackFunctionResponse); err != nil {
		return nil, err
	}
	return trackUntrackFunctionResponse, nil
}

func NewTrackUntrackFunctionQuery(untrack bool) Query {
	trackUntrackType := TRACK_FUNCTION_TYPE
	if untrack {
		trackUntrackType = UNTRACK_FUNCTION_TYPE
	}
	query := TrackUntrackFunction{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: trackUntrackType,
	}

	return Query(&query)
}
