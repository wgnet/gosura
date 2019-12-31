package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DELETE_EVENT_TRIGGER_TYPE string = `delete_event_trigger`
)

type DeleteEventTrigger struct {
	Arguments DeleteEventTriggerArgs `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type DeleteEventTriggerArgs struct {
	Name string `json:"name"`
}

type DeleteEventTriggerResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *DeleteEventTrigger) SetArgs(args interface{}) error {
	switch args.(type) {
	case DeleteEventTriggerArgs:
		t.Arguments = args.(DeleteEventTriggerArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *DeleteEventTrigger) SetVersion(version int) {
	t.Ver = version
}

func (t *DeleteEventTrigger) SetType(name string) {
	t.QueryType = name
}

func (t *DeleteEventTrigger) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *DeleteEventTrigger) Method() string {
	return http.MethodPost
}

func (t *DeleteEventTrigger) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var deleteEventTriggerResponse DeleteEventTriggerResponse
	if err := json.Unmarshal(body, &deleteEventTriggerResponse); err != nil {
		return nil, err
	}
	return deleteEventTriggerResponse, nil
}

func NewDeleteEventTriggerQuery() Query {
	query := DeleteEventTrigger{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: DELETE_EVENT_TRIGGER_TYPE,
	}

	return Query(&query)
}
