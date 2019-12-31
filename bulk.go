package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BULK_TYPE string = `bulk`
)

type Bulk struct {
	Arguments []Query `json:"args"`
	QueryType string  `json:"type"`
}

type BulkResponse []map[string]interface{}

// SetArgs adds a new Query to the Arguments field
func (t *Bulk) SetArgs(args interface{}) error {
	switch args.(type) {
	case Query:
		t.Arguments = append(t.Arguments, args.(Query))
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *Bulk) SetVersion(version int) {}

func (t *Bulk) SetType(name string) {
	t.QueryType = name
}

func (t *Bulk) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Bulk) Method() string {
	return http.MethodPost
}

func (t *Bulk) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var bulkResponse BulkResponse
	if err := json.Unmarshal(body, &bulkResponse); err != nil {
		return nil, err
	}
	return bulkResponse, nil
}

func NewBulkQuery() Query {
	query := Bulk{
		QueryType: BULK_TYPE,
	}

	return Query(&query)
}
