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
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

var ErrNotfoundClientIP = fmt.Errorf("not found client ip")
var ErrDomainParse = fmt.Errorf("domain parse failed")

// DomainParsed parsed domain result structure definition.
type DomainParsed struct {
	Original string `json:"original"`
	IPS      IPList `json:"ips"`
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

// IPv4s returns an IPv4 list.
func (d *DomainParsed) IPv4s() IPList {
	var list IPList
	for _, cur := range d.IPS {
		if IsIPv4(cur) {
			if list == nil {
				list = make(IPList, 0)
			}

			list = append(list, cur)
		}
	}

	return list
}

// IPv6s returns an IPv6 list.
func (d *DomainParsed) IPv6s() IPList {
	var list IPList
	for _, cur := range d.IPS {
		if IsIPv6(cur) {
			if list == nil {
				list = make(IPList, 0)
			}

			list = append(list, cur)
		}
	}

	return list
}

// NewDomainParsedWithOriginal returns a new DomainParsed instance with original
// domain string.
func NewDomainParsedWithOriginal(domain string) *DomainParsed {
	return &DomainParsed{
		Original: domain,
		IPS:      make(IPList, 0),
	}
}

// DomainParse parse domain to IP address.
func DomainParse(domain string) (*DomainParsed, error) {
	var parsed = NewDomainParsedWithOriginal(domain)

	// Quick lookup host.
	schema := strings.Contains(domain, "http")
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
			return nil, ErrDomainParse
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

func ClientIP(req *http.Request) (string, error) {
	// Get real IP.
	rea := req.Header.Get("X-Real-IP")
	if net.ParseIP(rea) != nil {
		return rea, nil
	}

	// Get proxy IP.
	proxy := req.Header.Get("X-Forwarded-IP")
	if net.ParseIP(proxy) != nil {
		return proxy, nil
	}

	return "", ErrNotfoundClientIP
}
