package Flags

import "errors"

type Flags struct {
	Flag map[string]string
}

var ErrUnsetFlags error

func init() {
	ErrUnsetFlags = errors.New("unset Flags")
}

func (flags Flags) CheckIfFlagsAreSet(wanted ...string) bool {
	for _, wantedFlag := range wanted {
		if flags.Flag[wantedFlag] == "" {
			return false
		}
	}
	return true
}
