package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	CREATE_OBJECT_RELATIONSHIP_TYPE string = `create_object_relationship`
)

type CreateObjectRelationship struct {
	Arguments CreateObjectRelationshipArgs `json:"args"`
	Ver       int                          `json:"version"`
	QueryType string                       `json:"type"`
}

type CreateObjectRelationshipArgs struct {
	Table   string      `json:"table"`
	Name    string      `json:"name"`
	Using   interface{} `json:"using"`
	Comment string      `json:"comment"`
}

type CreateObjectRelationshipUsingAuto struct {
	ForeignKeyConstraintOn string `json:"foreign_key_constraint_on"`
}

type CreateObjectRelationshipUsingManual struct {
	ManualConfiguration RelationshipManualConf `json:"manual_configuration"`
}

type RelationshipManualConf struct {
	RemoteTable   string            `json:"remote_table"`
	ColumnMapping map[string]string `json:"column_mapping"`
}

type CreateObjectRelationshipResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *CreateObjectRelationship) SetArgs(args interface{}) error {
	switch args.(type) {
	case CreateObjectRelationshipArgs:
		t.Arguments = args.(CreateObjectRelationshipArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *CreateObjectRelationship) SetVersion(version int) {
	t.Ver = version
}

func (t *CreateObjectRelationship) SetType(name string) {
	t.QueryType = name
}

func (t *CreateObjectRelationship) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *CreateObjectRelationship) Method() string {
	return http.MethodPost
}

func (t *CreateObjectRelationship) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var createObjectRelationshipResponse CreateObjectRelationshipResponse
	if err := json.Unmarshal(body, &createObjectRelationshipResponse); err != nil {
		return nil, err
	}
	return createObjectRelationshipResponse, nil
}

func NewCreateObjectRelationshipQuery() Query {
	query := CreateObjectRelationship{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: CREATE_OBJECT_RELATIONSHIP_TYPE,
	}

	return Query(&query)
}
