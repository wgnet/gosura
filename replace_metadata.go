package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	REPLACE_METADATA_TYPE string = `replace_metadata`
)

type ReplaceMetadata struct {
	Arguments ReplaceMetadataArgs `json:"args"`
	Ver       int                 `json:"version"`
	QueryType string              `json:"type"`
}

type ReplaceMetadataArgs map[string]interface{}
type ReplaceMetadataResponse map[string]interface{}

func (t *ReplaceMetadata) SetArgs(args interface{}) error {
	switch args.(type) {
	case ReplaceMetadataArgs:
		t.Arguments = args.(ReplaceMetadataArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *ReplaceMetadata) SetVersion(version int) {
	t.Ver = version
}

func (t *ReplaceMetadata) SetType(name string) {
	t.QueryType = name
}

func (t *ReplaceMetadata) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *ReplaceMetadata) Method() string {
	return http.MethodPost
}

func (t *ReplaceMetadata) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var replaceMetadataResponse ReplaceMetadataResponse
	if err := json.Unmarshal(body, &replaceMetadataResponse); err != nil {
		return nil, err
	}
	return replaceMetadataResponse, nil
}

func NewReplaceMetadataQuery() Query {
	query := ReplaceMetadata{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: REPLACE_METADATA_TYPE,
	}

	return Query(&query)
}
