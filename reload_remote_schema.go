package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	RELOAD_REMOTE_SCHEMA_TYPE string = `reload_remote_schema`
)

type ReloadRemoteSchema struct {
	Arguments ReloadRemoteSchemaArgs `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type ReloadRemoteSchemaArgs struct {
	Name string `json:"name"`
}

type ReloadRemoteSchemaResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *ReloadRemoteSchema) SetArgs(args interface{}) error {
	switch args.(type) {
	case ReloadRemoteSchemaArgs:
		t.Arguments = args.(ReloadRemoteSchemaArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *ReloadRemoteSchema) SetVersion(version int) {
	t.Ver = version
}

func (t *ReloadRemoteSchema) SetType(name string) {
	t.QueryType = name
}

func (t *ReloadRemoteSchema) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *ReloadRemoteSchema) Method() string {
	return http.MethodPost
}

func (t *ReloadRemoteSchema) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var reloadRemoteSchemaResponse ReloadRemoteSchemaResponse
	if err := json.Unmarshal(body, &reloadRemoteSchemaResponse); err != nil {
		return nil, err
	}
	return reloadRemoteSchemaResponse, nil
}

func NewReloadRemoteSchemaQuery() Query {
	query := ReloadRemoteSchema{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: RELOAD_REMOTE_SCHEMA_TYPE,
	}

	return Query(&query)
}
