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
	"io/ioutil"
	"net/http"
)

const (
	DEFAULT_CONFIG_ENDPOINT_PATH string = `/v1alpha1/config`
)

type Config struct{}

type ConfigResponse map[string]interface{}

func (c *Config) SetArgs(args interface{}) error {
	return nil
}

func (c *Config) SetVersion(version int) {}

func (r *Config) SetType(name string) {}

func (r *Config) Byte() ([]byte, error) {
	return []byte(""), nil
}

func (r *Config) Method() string {
	return http.MethodGet
}

func (r *Config) CheckResponse(response *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode > http.StatusOK {
		return body, fmt.Errorf("Error received")
	}

	var configResponse ConfigResponse
	if err := json.Unmarshal(body, &configResponse); err != nil {
		return nil, err
	}
	return configResponse, nil
}

func NewConfigQuery() Query {
	query := Config{}

	return Query(&query)
}
