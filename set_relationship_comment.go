package gosura

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SET_RELATIONSHIP_COMMENT_TYPE string = `set_relationship_comment`
)

type SetRelationshipComment struct {
	Arguments SetRelationshipCommentArgs `json:"args"`
	Ver       int                        `json:"version"`
	QueryType string                     `json:"type"`
}

type SetRelationshipCommentArgs struct {
	Table   string `json:"table"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type SetRelationshipCommentResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

func (t *SetRelationshipComment) SetArgs(args interface{}) error {
	switch args.(type) {
	case SetRelationshipCommentArgs:
		t.Arguments = args.(SetRelationshipCommentArgs)
	default:
		return fmt.Errorf("Wrong args type %T", args)
	}
	return nil
}

func (t *SetRelationshipComment) SetVersion(version int) {
	t.Ver = version
}

func (t *SetRelationshipComment) SetType(name string) {
	t.QueryType = name
}

func (t *SetRelationshipComment) Byte() ([]byte, error) {
	return json.Marshal(t)
}

func (t *SetRelationshipComment) Method() string {
	return http.MethodPost
}

func (t *SetRelationshipComment) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := checkResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var setRelationshipResponse SetRelationshipCommentResponse
	if err := json.Unmarshal(body, &setRelationshipResponse); err != nil {
		return nil, err
	}
	return setRelationshipResponse, nil
}

func NewSetRelationshipCommentQuery() Query {
	query := SetRelationshipComment{
		Ver:       DEFAULT_QUERY_VERSION,
		QueryType: SET_RELATIONSHIP_COMMENT_TYPE,
	}

	return Query(&query)
}
