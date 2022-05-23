// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>
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

package netio

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// DomainParsed parsed domain result structure definition.
type DomainParsed struct {
	Original string   `json:"original"`
	IPS      []string `json:"ips"`
}

// FirstIP returns the IPS list first IP address.
func (d *DomainParsed) FirstIP() string {
	if len(d.IPS) > 0 {
		return d.IPS[0]
	}

	return ""
}

// String implement Stringer interface.
func (d *DomainParsed) String() string {
	data, _ := json.Marshal(&d)
	return string(data)
}

// NewDomainParsedWithOriginal returns a new DomainParsed instance with original
// domain string.
func NewDomainParsedWithOriginal(domain string) *DomainParsed {
	return &DomainParsed{
		Original: domain,
		IPS:      make([]string, 0),
	}
}

// DomainParse parse domain to IP address.
func DomainParse(domain string) (*DomainParsed, error) {
	var parsed = NewDomainParsedWithOriginal(domain)

	// Quick lookup host.
	schema := strings.ContainsAny(domain, "http")
	if !schema {
		host, err := net.LookupHost(domain)
		if err != nil {
			return nil, err
		}

		parsed.IPS = append(parsed.IPS, host...)

		return parsed, nil
	}

	// Get IP address by protocol type.
	req, err := http.NewRequest(http.MethodGet, domain, nil)
	if err != nil {
		return nil, err
	}

	dail := func(ctx context.Context, network, addr string) (net.Conn, error) {
		dial, err := net.Dial(network, addr)
		if err != nil {
			return nil, err
		}

		slice := strings.Split(dial.RemoteAddr().String(), ":")
		if len(slice) == 0 {
			return nil, fmt.Errorf("domain parse failed")
		}

		parsed.IPS = append(parsed.IPS, slice[0])

		return dial, nil
	}

	client := &http.Client{Transport: &http.Transport{DialContext: dail}}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return parsed, nil
}
