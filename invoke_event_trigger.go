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
