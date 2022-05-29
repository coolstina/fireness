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
	"net"
)

// IsIPv4 Check whether the IP address is IPv4.
func IsIPv4(addr string) bool {
	if !IsIP(addr) {
		return false
	}

	ip := net.ParseIP(addr)
	if ip.To4() != nil {
		return true
	}

	return false
}

// IsIPv6 Check whether the IP address is IPv6.
func IsIPv6(addr string) bool {
	if !IsIP(addr) {
		return false
	}

	ip := net.ParseIP(addr)
	if ip.To4() != nil {
		return false
	}
	return true
}

// IsIP Check whether the IP address is real ip.
func IsIP(addr string) bool {
	ip := net.ParseIP(addr)
	if ip != nil {
		return true
	}

	return false
}
