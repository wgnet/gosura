package gosura

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DEFAULT_ENDPOINT_PATH string = `/v1/query`
	DEFAULT_QUERY_VERSION int    = 1
	QUERY_VERSION_V2      int    = 2
)

type Query interface {
	SetType(t string)
	SetArgs(args interface{}) error
	SetVersion(version int)
	Method() string
	Byte() ([]byte, error)
	CheckResponse(response *http.Response, err error) (interface{}, error)
}

type HasuraQueryError struct {
	Path  string `json:"path"`
	Error string `json:"error"`
}

func checkResponseStatus(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return body, fmt.Errorf("Can't read response body: %w", err)
	}

	if response.StatusCode >= http.StatusBadRequest {
		var queryErr HasuraQueryError
		if err := json.Unmarshal(body, &queryErr); err != nil {
			return body, fmt.Errorf("Can't unmarshal error response: %w", err)
		}
		return body, fmt.Errorf("Error received in path %s: %s", queryErr.Path, queryErr.Error)
	}

	return body, nil
}
