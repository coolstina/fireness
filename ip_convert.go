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
	"encoding/binary"
	"fmt"
	"net"
)

// IPv4ToInt Convert IPv4 address string to integer.
func IPv4ToInt(ip string) uint32 {
	return binary.BigEndian.Uint32(net.ParseIP(ip).To4())
}

// IntToIPv4 Convert IP integer to IPv4 address string.
func IntToIPv4(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
