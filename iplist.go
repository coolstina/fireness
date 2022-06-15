package fireness

// IPList IP list structure definition.
type IPList []string

// IsEmpty detect IP list is empty.
func (iplist IPList) IsEmpty() bool {
	if len(iplist) == 0 {
		return true
	}

	return false
}

// IsNil detect IP list is nil.
func (iplist IPList) IsNil() bool {
	if iplist == nil {
		return true
	}

	return false
}
