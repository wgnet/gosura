/**
 * Copyright 2019-2020 Wargaming Group Limited
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
**/

package gosura

import (
	"encoding/json"
)

type SelectColumn struct {
	name    string
	columns []*SelectColumn
}

func (s *SelectColumn) MarshalJSON() ([]byte, error) {
	if len(s.columns) == 0 {
		return json.Marshal(s.name)
	}
	column := make(map[string]interface{})
	column["name"] = s.name
	column["columns"] = s.columns
	return json.Marshal(column)
}

func (s *SelectColumn) AddColumn(name string, columns []*SelectColumn) *SelectColumn {
	s.columns = append(s.columns, &SelectColumn{name: name, columns: columns})
	return s
}

func NewSelectColumn(name string) *SelectColumn {
	return &SelectColumn{
		name: name,
	}
}
