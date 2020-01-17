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
	INVOKE_EVENT_TRIGGER_TYPE string = `invoke_event_trigger`
)

type InvokeEventTrigger struct {
	Arguments InvokeEventTriggerArgs `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type InvokeEventTriggerArgs struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"` // Payload cannot be empty
}

func NewInvokeEventTriggerArgs(name string, payload interface{}) InvokeEventTriggerArgs {
	args := InvokeEventTriggerArgs{name, payload}
	if payload == nil {
		emptyPayload := make(map[string]interface{})
		args.Payload = emptyPayload
	}
	return args
}

type InvokeEventTriggerResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *InvokeEventTrigger) SetArgs(args interface{}) error {
	switch args.(type) {
	case InvokeEventTriggerArgs:
		t.Arguments = args.(InvokeEventTriggerArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *InvokeEventTrigger) SetVersion(version int) {
	t.Ver = version
}

func (t *InvokeEventTrigger) SetType(name string) {
	t.QueryType = name
}

func (t *InvokeEventTrigger) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *InvokeEventTrigger) Method() string {
	return http.MethodPost
}

func (t *InvokeEventTrigger) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var invokeEventTriggerResponse InvokeEventTriggerResponse
	if err := json.Unmarshal(body, &invokeEventTriggerResponse); err != nil {
		return nil, err
	}
	return invokeEventTriggerResponse, nil
}

func NewInvokeEventTriggerQuery() Query {
	query := InvokeEventTrigger{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: INVOKE_EVENT_TRIGGER_TYPE,
	}

	return Query(&query)
}
