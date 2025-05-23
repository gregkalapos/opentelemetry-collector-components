// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package configintegrationextension // import "github.com/elastic/opentelemetry-collector-components/extension/configintegrationextension"

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIntegrationValidation(t *testing.T) {
	cases := []struct {
		title       string
		configFile  string
		expectedErr string
	}{
		{
			title:      "valid",
			configFile: "testdata/valid-config-with-templates.yaml",
		},
		{
			title:       "invalid yaml format",
			configFile:  "testdata/config-with-templates.yaml",
			expectedErr: `could not find expected ':'`,
		},
		{
			title:       "missing component",
			configFile:  "testdata/config-unknown-key.yaml",
			expectedErr: `field extensions not found`,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			config := configFromFile(t, c.configFile)
			err := config.Validate()
			if c.expectedErr != "" {
				assert.ErrorContains(t, err, c.expectedErr)
				return
			}

			assert.NoError(t, err)
		})
	}
}
