package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	RUN_SQL_TYPE string = `run_sql`
)

type RunSql struct {
	Arguments RunSqlArgs `json:"args"`
	Ver       int        `json:"version"`
	QueryType string     `json:"type"`
}

type RunSqlArgs struct {
	SQL                      string `json:"sql"`
	Cascade                  bool   `json:"cascade"`
	CheckMetadataConsistency bool   `json:"check_metadata_consistency"`
}

type RunSqlResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (r *RunSql) SetArgs(args interface{}) error {
	switch args.(type) {
	case RunSqlArgs:
		r.Arguments = args.(RunSqlArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (r *RunSql) SetVersion(version int) {
	r.Ver = version
}

func (r *RunSql) SetType(name string) {
	r.QueryType = name
}

func (r *RunSql) Byte() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RunSql) Method() string {
	return http.MethodPost
}

func (r *RunSql) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var runSqlResponse RunSqlResponse
	if err := json.Unmarshal(body, &runSqlResponse); err != nil {
		return nil, err
	}
	return runSqlResponse, nil
}

func NewRunSqlQuery() Query {
	query := RunSql{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: RUN_SQL_TYPE,
	}

	return Query(&query)
}
