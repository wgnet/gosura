package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DROP_INSERT_PERMISSION_TYPE string = `drop_insert_permission`
	DROP_SELECT_PERMISSION_TYPE string = `drop_select_permission`
	DROP_UPDATE_PERMISSION_TYPE string = `drop_update_permission`
	DROP_DELETE_PERMISSION_TYPE string = `drop_delete_permission`
)

type DropPermission struct {
	Arguments DropPermissionArgs `json:"args"`
	QueryType string             `json:"type"`
}

type DropPermissionArgs struct {
	Table string `json:"table"`
	Role  string `json:"role"`
}

type DropPermissionResponse map[string]interface{}

func (t *DropPermission) SetArgs(args interface{}) error {
	switch args.(type) {
	case DropPermissionArgs:
		t.Arguments = args.(DropPermissionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *DropPermission) SetVersion(_ int) {}

func (t *DropPermission) SetType(name string) {
	t.QueryType = name
}

func (t *DropPermission) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *DropPermission) Method() string {
	return http.MethodPost
}

func (t *DropPermission) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var dropPermissionResponse DropPermissionResponse
	if err := json.Unmarshal(body, &dropPermissionResponse); err != nil {
		return nil, err
	}
	return dropPermissionResponse, nil
}

func NewDropPermissionQuery(t string) Query {
	var dropType string
	switch t {
	case DROP_INSERT_PERMISSION_TYPE:
		dropType = DROP_INSERT_PERMISSION_TYPE
	case DROP_SELECT_PERMISSION_TYPE:
		dropType = DROP_SELECT_PERMISSION_TYPE
	case DROP_UPDATE_PERMISSION_TYPE:
		dropType = DROP_UPDATE_PERMISSION_TYPE
	case DROP_DELETE_PERMISSION_TYPE:
		dropType = DROP_DELETE_PERMISSION_TYPE
	}
	query := DropPermission{
		QueryType: dropType,
	}

	return Query(&query)
}
