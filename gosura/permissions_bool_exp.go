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
	"fmt"
	"sync"
)

const (
	AND_EXP_TYPE    string = `$and`
	OR_EXP_TYPE     string = `$or`
	NOT_EXP_TYPE    string = `$not`
	EXISTS_EXP_TYPE string = `$exists`
	TRUE_EXP_TYPE   string = `true`

	// PG_OPERATORS
	PG_OPERATOR_EQ       string = `$eq`
	PG_OPERATOR_NE       string = `$ne`
	PG_OPERATOR_IN       string = `$in`
	PG_OPERATOR_NIN      string = `$nin`
	PG_OPERATOR_GT       string = `$gt`
	PG_OPERATOR_LT       string = `$lt`
	PG_OPERATOR_GTE      string = `$gte`
	PG_OPERATOR_LTE      string = `$lte`
	PG_OPERATOR_LIKE     string = `$like`
	PG_OPERATOR_NLIKE    string = `$nlike`
	PG_OPERATOR_ILIKE    string = `$ilike`
	PG_OPERATOR_NILIKE   string = `$nilike`
	PG_OPERATOR_SIMILAR  string = `$similar`
	PG_OPERATOR_NSIMILAR string = `$nsimilar`
	PG_OPERATOR_CEQ      string = `$ceq`
	PG_OPERATOR_CNE      string = `$cne`
	PG_OPERATOR_CGT      string = `$cgt`
	PG_OPERATOR_CLT      string = `$clt`
	PG_OPERATOR_CGTE     string = `$cgte`
	PG_OPERATOR_CLTE     string = `$clte`

	PG_OPERATOR_IS_NULL string = "_is_null"

	PG_JSONB_OPERATOR_CONTAINS     string = `_contains`
	PG_JSONB_OPERATOR_CONTAINED_IN string = `_contained_in`
	PG_JSONB_OPERATOR_HAS_KEY      string = `_has_key`
	PG_JSONB_OPERATOR_HAS_KEYS_ANY string = `_has_keys_any`
	PG_JSONB_OPERATOR_HAS_KEYS_ALL string = `_has_keys_all`

	PG_GIS_OPERATOR_CONTAINS   string = `_st_contains`
	PG_GIS_OPERATOR_CROSSES    string = `_st_crosses`
	PG_GIS_OPERATOR_EQUALS     string = `_st_equals`
	PG_GIS_OPERATOR_INTERSECTS string = `_st_intersects`
	PG_GIS_OPERATOR_OVERLAPS   string = `_st_overlaps`
	PG_GIS_OPERATOR_TOUCHES    string = `_st_touches`
	PG_GIS_OPERATOR_WITHIN     string = `_st_within`
	PG_GIS_OPERATOR_D_WITHIN   string = `_st_d_within`
)

type BoolExp struct {
	mux sync.Mutex
	exp map[string]interface{}
}

func (b *BoolExp) addAndExp(exp ...Bool) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.exp == nil {
		b.exp = make(map[string]interface{})
	}
	if _, ok := b.exp[AND_EXP_TYPE]; !ok {
		b.exp[AND_EXP_TYPE] = []map[string]interface{}{}
	}

	for _, e := range exp {
		b.exp[AND_EXP_TYPE] = append(b.exp[AND_EXP_TYPE].([]map[string]interface{}), e.Map())
	}
}

func (b *BoolExp) addOrExp(exp ...Bool) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.exp == nil {
		b.exp = make(map[string]interface{})
	}
	if _, ok := b.exp[OR_EXP_TYPE]; !ok {
		b.exp[OR_EXP_TYPE] = []map[string]interface{}{}
	}

	for _, e := range exp {
		b.exp[OR_EXP_TYPE] = append(b.exp[OR_EXP_TYPE].([]map[string]interface{}), e.Map())
	}
}

func (b *BoolExp) addNotExp(exp Bool) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.exp == nil {
		b.exp = make(map[string]interface{})
	}

	b.exp[NOT_EXP_TYPE] = exp.Map()
}

func (b *BoolExp) addExistsExp(exp Bool) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.exp == nil {
		b.exp = make(map[string]interface{})
	}

	b.exp[EXISTS_EXP_TYPE] = exp.Map()
}

func (b *BoolExp) addTrueExp() {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.exp == nil {
		b.exp = make(map[string]interface{})
	}
}

func (b *BoolExp) copy() map[string]interface{} {
	b.mux.Lock()
	defer b.mux.Unlock()

	newCopy := make(map[string]interface{})
	for k, v := range b.exp {
		newCopy[k] = v
	}
	return newCopy
}

func (b *BoolExp) Map() map[string]interface{} {
	return b.copy()
}

func (b *BoolExp) AddExp(expType string, exp ...Bool) error {
	switch expType {
	case AND_EXP_TYPE:
		b.addAndExp(exp...)
		return nil
	case OR_EXP_TYPE:
		b.addOrExp(exp...)
		return nil
	case NOT_EXP_TYPE:
		if len(exp) > 0 {
			b.addNotExp(exp[0])
		}
		return nil
	case EXISTS_EXP_TYPE:
		if len(exp) > 0 {
			b.addExistsExp(exp[0])
		}
	case TRUE_EXP_TYPE:
		b.addTrueExp()
	default:
		return fmt.Errorf("Wrong bool expression type %s", expType)
	}
	return nil
}

func (b *BoolExp) AddKV(key string, value interface{}) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.exp == nil {
		b.exp = make(map[string]interface{})
	}
	b.exp[key] = value
}

func (b *BoolExp) SetRaw(check map[string]interface{}) {
	b.mux.Lock()
	b.mux.Unlock()
	b.exp = check
}

func (b *BoolExp) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Map())
}

func NewBoolExp() Bool {
	return Bool(&BoolExp{
		exp: make(map[string]interface{}),
	})
}

type Bool interface {
	AddExp(expType string, exp ...Bool) error
	AddKV(key string, value interface{})
	Map() map[string]interface{}
	MarshalJSON() ([]byte, error)
	SetRaw(check map[string]interface{})
}
