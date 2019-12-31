package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SET_TABLE_CUSTOM_FIELDS_TYPE string = `set_table_custom_fields`
)

type SetTableCustomFields struct {
	Arguments SetTableCustomFieldsArgs `json:"args"`
	Ver       int                      `json:"version"`
	QueryType string                   `json:"type"`
}

type SetTableCustomFieldsArgs struct {
	Table             string            `json:"table"`
	CustomRootFields  CustomRootFields  `json:"custom_root_fields"`
	CustomColumnNames map[string]string `json:"custom_column_names"`
}

type SetTableCustomFieldsResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *SetTableCustomFields) SetArgs(args interface{}) error {
	switch args.(type) {
	case SetTableCustomFieldsArgs:
		t.Arguments = args.(SetTableCustomFieldsArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *SetTableCustomFields) SetVersion(version int) {
	t.Ver = version
}

func (t *SetTableCustomFields) SetType(name string) {
	t.QueryType = name
}

func (t *SetTableCustomFields) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *SetTableCustomFields) Method() string {
	return http.MethodPost
}

func (t *SetTableCustomFields) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var trackTableResponse SetTableCustomFieldsResponse
	if err := json.Unmarshal(body, &trackTableResponse); err != nil {
		return nil, err
	}
	return trackTableResponse, nil
}

func NewSetTableCustomFieldsQuery() Query {
	query := SetTableCustomFields{
		Ver:       QUERY_VERSION_V2,
		QueryType: SET_TABLE_CUSTOM_FIELDS_TYPE,
	}

	return Query(&query)
}
