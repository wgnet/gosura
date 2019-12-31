package gosura

import (
	"encoding/json"
	"net/http"
)

const (
	CLEAR_METADATA_TYPE string = `clear_metadata`
)

type ClearMetadata struct {
	Arguments map[string]interface{} `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type ClearMetadataResponse map[string]interface{}

// SetArgs do nothing here
func (t *ClearMetadata) SetArgs(args interface{}) error {
	return nil
}

func (t *ClearMetadata) SetVersion(version int) {
	t.Ver = version
}

func (t *ClearMetadata) SetType(name string) {
	t.QueryType = name
}

func (t *ClearMetadata) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *ClearMetadata) Method() string {
	return http.MethodPost
}

func (t *ClearMetadata) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var clearMetadataResponse ClearMetadataResponse
	if err := json.Unmarshal(body, &clearMetadataResponse); err != nil {
		return nil, err
	}
	return clearMetadataResponse, nil
}

func NewClearMetadataQuery() Query {
	query := ClearMetadata{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: CLEAR_METADATA_TYPE,
		Arguments: make(map[string]interface{}),
	}

	return Query(&query)
}
