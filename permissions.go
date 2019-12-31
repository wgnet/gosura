package gosura

import "encoding/json"

type PGColumn struct {
	all   string
	array []string
}

func (p *PGColumn) MarshalJSON() ([]byte, error) {
	if p.all == "*" {
		return json.Marshal(p.all)
	}
	return json.Marshal(p.array)
}

func (p *PGColumn) AddColumn(col ...string) *PGColumn {
	if len(col) == 1 && col[0] == "*" {
		p.all = "*"
		return p
	}
	for i, _ := range col {
		p.array = append(p.array, col[i])
	}
	return p
}

func NewPGColumn() *PGColumn {
	return &PGColumn{}
}

type InsertPermission struct {
	Check   Bool                   `json:"check"`
	Set     map[string]interface{} `json:"set,omitempty"`
	Columns *PGColumn              `json:"columns"`
}

func NewInsertPermission() *InsertPermission {
	return &InsertPermission{}
}

type SelectPermission struct {
	Columns           *PGColumn `json:"columns"`
	Filter            Bool      `json:"filter"`
	ComputedFields    []string  `json:"computed_fields,omitempty"`
	Limit             int64     `json:"limit,omitempty"`
	AllowAggregations bool      `json:"allow_aggregations,omitempty"`
}

func NewSelectPermission() *SelectPermission {
	return &SelectPermission{}
}

type UpdatePermission struct {
	Columns *PGColumn              `json:"columns"`
	Filter  Bool                   `json:"filter"`
	Set     map[string]interface{} `json:"set,omitempty"`
}

func NewUpdatePermission() *UpdatePermission {
	return &UpdatePermission{}
}

type DeletePermission struct {
	Filter Bool `json:"filter"`
}

func NewDeletePermission() *DeletePermission {
	return &DeletePermission{}
}
