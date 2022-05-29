// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fireness

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFirebaseSuite(t *testing.T) {
	suite.Run(t, &FirebaseSuite{})
}

type FirebaseSuite struct {
	suite.Suite
}

func (suite *FirebaseSuite) Test_DomainParse() {
	grids := []struct {
		domain          string
		hasError        bool
		hasErrorString  string
		expectedString  string
		expectedCompare bool
	}{
		// {"original":"baidu.com","ips":["220.181.38.148","220.181.38.251"]}
		{
			domain: "baidu.com",
		},
		// {"original":"https://baidu.com","ips":["220.181.38.148","103.235.46.39"]}
		{
			domain: "https://baidu.com",
		},
		{
			domain:         "localhost",
			hasError:       true,
			hasErrorString: `unsupported protocol scheme ""`,
		},
		{
			domain:          "127.0.0.1",
			expectedString:  `{"original":"127.0.0.1","ips":["127.0.0.1"]}`,
			expectedCompare: true,
		},
		{
			domain:          "220.181.38.148",
			expectedString:  `{"original":"220.181.38.148","ips":["220.181.38.148"]}`,
			expectedCompare: true,
		},
		{
			domain:          "mrbs.autosolan.com",
			expectedString:  `{"original":"mrbs.autosolan.com","ips":["113.240.251.7"]}`,
			expectedCompare: true,
		},
	}

	for _, grid := range grids {
		parse, err := DomainParse(grid.domain)
		if err != nil {
			assert.Contains(suite.T(), err.Error(), grid.hasErrorString)
		} else {
			assert.NoError(suite.T(), err)
			if grid.expectedCompare {
				assert.Equal(suite.T(), grid.expectedString, parse.String())
			}
		}

		log.Printf("parse: %+v\n", parse)
	}
}

func (suite *FirebaseSuite) Test_ClientIP() {
	request := httptest.NewRequest("GET", "/", nil)
	request.Header.Set("x-real-ip", "192.168.0.1")
	request.Header.Set("x-forwarded-for", "192.168.0.255")

	actual, err := ClientIP(request)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), `192.168.0.1`, actual)
}
