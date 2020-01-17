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
	CREATE_EVENT_TRIGGER_TYPE string = `create_event_trigger`
)

type CreateEventTrigger struct {
	Arguments CreateEventTriggerArgs `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type CreateEventTriggerArgs struct {
	Name         string               `json:"name"`
	Table        string               `json:"table"`
	Webhook      string               `json:"webhook"`
	Insert       *EventTriggerPayload `json:"insert,omitempty"`
	Update       *EventTriggerPayload `json:"update,omitempty"`
	Delete       *EventTriggerPayload `json:"delete,omitempty"`
	Headers      []map[string]string  `json:"headers,omitempty"`
	Replace      bool                 `json:"replace"`
	EnableManual bool                 `json:"enable_manual"`
}

type EventTriggerPayload struct {
	Column  interface{} `json:"columns,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

// Payloads add payloads to the Payload field
// makes an slice if len of columns is larger than 0
// makes a string with `*` if columns is empty
func (e *EventTriggerPayload) Payloads(payloads ...string) {
	if len(payloads) > 0 {
		var p []string
		for i, _ := range payloads {
			p = append(p, payloads[i])
		}
		e.Payload = p
		return
	}

	e.Payload = "*"
}

// Columns add columns to the Columns field
// makes an slice if len of columns is larger than 0
// makes a string with `*` if columns is empty
func (e *EventTriggerPayload) Columns(columns ...string) {
	if len(columns) > 0 {
		var cols []string
		for i, _ := range columns {
			cols = append(cols, columns[i])
		}
		e.Column = cols
		return
	}

	e.Column = "*"
}

type CreateEventTriggerUsingAuto struct {
	ForeignKeyConstraintOn CreateEventTriggerUsingAutoConstraint `json:"foreign_key_constraint_on"`
}

type CreateEventTriggerUsingAutoConstraint struct {
	Table  string `json:"table"`
	Column string `json:"column"`
}

type CreateEventTriggerResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *CreateEventTrigger) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateEventTriggerArgs:
		t.Arguments = args.(CreateEventTriggerArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateEventTrigger) SetVersion(version int) {
	t.Ver = version
}

func (t *CreateEventTrigger) SetType(name string) {
	t.QueryType = name
}

func (t *CreateEventTrigger) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateEventTrigger) Method() string {
	return http.MethodPost
}

func (t *CreateEventTrigger) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createEventTriggerResponse CreateEventTriggerResponse
	if err := json.Unmarshal(body, &createEventTriggerResponse); err != nil {
		return nil, err
	}
	return createEventTriggerResponse, nil
}

func NewCreateEventTriggerQuery() Query {
	query := CreateEventTrigger{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: CREATE_EVENT_TRIGGER_TYPE,
	}

	return Query(&query)
}
