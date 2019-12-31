package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ADD_QUERY_TO_COLLECTION_TYPE string = `add_query_to_collection`
)

type AddQueryToCollection struct {
	Arguments AddQueryToCollectionArgs `json:"args"`
	Ver       int                      `json:"version"`
	QueryType string                   `json:"type"`
}

type AddQueryToCollectionArgs struct {
	Collection string `json:"collection_name"`
	Name       string `json:"query_name"`
	Query      string `json:"query"`
}

type AddQueryToCollectionResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *AddQueryToCollection) SetArgs(args interface{}) error {
	switch args.(type) {
	case AddQueryToCollectionArgs:
		t.Arguments = args.(AddQueryToCollectionArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *AddQueryToCollection) SetVersion(version int) {
	t.Ver = version
}

func (t *AddQueryToCollection) SetType(name string) {
	t.QueryType = name
}

func (t *AddQueryToCollection) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *AddQueryToCollection) Method() string {
	return http.MethodPost
}

func (t *AddQueryToCollection) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var addQueryToCollectionResponse AddQueryToCollectionResponse
	if err := json.Unmarshal(body, &addQueryToCollectionResponse); err != nil {
		return nil, err
	}
	return addQueryToCollectionResponse, nil
}

func NewAddQueryToCollectionQuery() Query {
	query := AddQueryToCollection{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: ADD_QUERY_TO_COLLECTION_TYPE,
	}

	return Query(&query)
}
