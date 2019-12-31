package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ADD_COLLECTION_TO_ALLOWLIST_TYPE string = `add_collection_to_allowlist`
)

type AddCollectionToAllowlist struct {
	Arguments AddCollectionToAllowlistArgs `json:"args"`
	Ver       int                          `json:"version"`
	QueryType string                       `json:"type"`
}

type AddCollectionToAllowlistArgs struct {
	Collection string `json:"collection"`
}

type AddCollectionToAllowlistResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *AddCollectionToAllowlist) SetArgs(args interface{}) error {
	switch args.(type) {
	case AddCollectionToAllowlistArgs:
		t.Arguments = args.(AddCollectionToAllowlistArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *AddCollectionToAllowlist) SetVersion(version int) {
	t.Ver = version
}

func (t *AddCollectionToAllowlist) SetType(name string) {
	t.QueryType = name
}

func (t *AddCollectionToAllowlist) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *AddCollectionToAllowlist) Method() string {
	return http.MethodPost
}

func (t *AddCollectionToAllowlist) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var dropQueryFromCollectionResponse AddCollectionToAllowlistResponse
	if err := json.Unmarshal(body, &dropQueryFromCollectionResponse); err != nil {
		return nil, err
	}
	return dropQueryFromCollectionResponse, nil
}

func NewAddCollectionToAllowlistQuery() Query {
	query := AddCollectionToAllowlist{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: ADD_COLLECTION_TO_ALLOWLIST_TYPE,
	}

	return Query(&query)
}
