package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	TRACK_TABLE_TYPE string = `track_table`
)

type TrackTable struct {
	Arguments TrackTableArgs `json:"args"`
	Ver       int            `json:"version"`
	QueryType string         `json:"type"`
}

type TrackTableArgs struct {
	Table  TableArgs `json:"table"`
	IsEnum bool      `json:"is_enum"`
}

type TrackTableResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *TrackTable) SetArgs(args interface{}) error {
	switch args.(type) {
	case TrackTableArgs:
		t.Arguments = args.(TrackTableArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *TrackTable) SetVersion(version int) {
	t.Ver = version
}

func (t *TrackTable) SetType(name string) {
	t.QueryType = name
}

func (t *TrackTable) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TrackTable) Method() string {
	return http.MethodPost
}

func (t *TrackTable) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var trackTableResponse TrackTableResponse
	if err := json.Unmarshal(body, &trackTableResponse); err != nil {
		return nil, err
	}
	return trackTableResponse, nil
}

func NewTrackTableQuery() Query {
	query := TrackTable{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: TRACK_TABLE_TYPE,
	}

	return Query(&query)
}
