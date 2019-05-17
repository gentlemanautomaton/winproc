package ntstatus

import "strconv"

// Value holds an NTSTATUS value.
type Value uint32

// String returns a string representation of the value.
func (v Value) String() string {
	if text, ok := descriptions[v]; ok {
		return text
	}
	return "NTSTATUS(" + strconv.FormatInt(int64(v), 16) + ")"
}

// Error returns an error string for the value.
func (v Value) Error() string {
	return v.String()
}
