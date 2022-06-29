package fireness

import (
	"net"
)

// LocalhostList localhost IP list structure definition.
type LocalhostList []string

// First returns localhost IP list first element.
func (list LocalhostList) First() string {
	if list != nil {
		return list[0]
	}

	return ""
}

// Localhost Get localhost IP list.
func Localhost(options ...LocalhostOption) (LocalhostList, error) {
	var ops = &localhostOption{}
	var list LocalhostList

	for _, o := range options {
		o.apply(ops)
	}

	faces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// handle err
	for _, i := range faces {
		adders, err := i.Addrs()
		if err != nil {
			continue
		}

		// handle err
		for _, addr := range adders {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			str := ip.String()
			if ops.excludes != nil {
				if inArray(ops.excludes, str) {
					continue
				}
			}

			if ops.excludeIPv4 {
				if IsIPv4(str) {
					continue
				}
			}

			if ops.excludeIPv6 {
				if IsIPv6(str) {
					continue
				}
			}

			if list == nil {
				list = make(LocalhostList, 0)
			}

			list = append(list, str)
		}
	}

	return list, nil
}

func inArray(data []string, find string) bool {
	for _, cur := range data {
		if find == cur {
			return true
		}
	}

	return false
}
