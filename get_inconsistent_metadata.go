package gosura

import (
	"encoding/json"
	"net/http"
)

const (
	GET_INCONSISTENT_METADATA_TYPE string = `get_inconsistent_metadata`
)

type GetInconsistentMetadata struct {
	Arguments map[string]interface{} `json:"args"`
	Ver       int                    `json:"version"`
	QueryType string                 `json:"type"`
}

type GetInconsistentMetadataResponse []map[string]interface{}

// SetArgs do nothing here
func (t *GetInconsistentMetadata) SetArgs(args interface{}) error {
	return nil
}

func (t *GetInconsistentMetadata) SetVersion(version int) {
	t.Ver = version
}

func (t *GetInconsistentMetadata) SetType(name string) {
	t.QueryType = name
}

func (t *GetInconsistentMetadata) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *GetInconsistentMetadata) Method() string {
	return http.MethodPost
}

func (t *GetInconsistentMetadata) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var getInconsistentMetadataResponse GetInconsistentMetadataResponse
	if err := json.Unmarshal(body, &getInconsistentMetadataResponse); err != nil {
		return nil, err
	}
	return getInconsistentMetadataResponse, nil
}

func NewGetInconsistentMetadataQuery() Query {
	query := GetInconsistentMetadata{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: GET_INCONSISTENT_METADATA_TYPE,
		Arguments: make(map[string]interface{}),
	}

	return Query(&query)
}
