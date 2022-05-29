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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIPv4(t *testing.T) {
	grids := []struct {
		ip       string
		expected bool
	}{
		{
			ip:       "hello world",
			expected: false,
		},
		{
			ip:       "中国北京",
			expected: false,
		},
		{
			ip:       "1000.40.210.253",
			expected: false,
		},
		{
			ip:       "127.0.0.1",
			expected: true,
		},
		{
			ip:       "::",
			expected: false,
		},
		{
			ip:       "::1",
			expected: false,
		},
		{
			ip:       "0:0:0:0:0:0",
			expected: false,
		},
		{
			ip:       "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := IsIPv4(grid.ip)
		assert.Equal(t, grid.expected, actual, "want: %+v, but got: %+v\n", grid.expected, actual)
	}
}

func TestIsIPv6(t *testing.T) {
	grids := []struct {
		ip       string
		expected bool
	}{
		{
			ip:       "::1",
			expected: true,
		},
		{
			ip:       "fe80::aede:48ff:fe00:1122",
			expected: true,
		},
		{
			ip:       "0.0.0.0",
			expected: false,
		},
		{
			ip:       "192.168.0.1",
			expected: false,
		},
		{
			ip:       "127.0.0.1",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := IsIPv6(grid.ip)
		assert.Equal(t, grid.expected, actual, "want: %+v, but got: %+v\n", grid.expected, actual)
	}
}

func TestIsIP(t *testing.T) {
	grids := []struct {
		ip       string
		expected bool
	}{
		{
			ip:       "hello world",
			expected: false,
		},
		{
			ip:       "中国北京",
			expected: false,
		},
		{
			ip:       "1000.40.210.253",
			expected: false,
		},
		{
			ip:       "127.0.0.1",
			expected: true,
		},
		{
			ip:       "::",
			expected: true,
		},
		{
			ip:       "::1",
			expected: true,
		},
		{
			ip:       "0:0:0:0:0:0",
			expected: false,
		},
		{
			ip:       "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := IsIP(grid.ip)
		assert.Equal(t, grid.expected, actual, "want: %+v, but got: %+v\n", grid.expected, actual)
	}
}
