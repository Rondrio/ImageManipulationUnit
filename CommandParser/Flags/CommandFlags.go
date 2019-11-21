package Flags

type Flags struct {
	Flag map[string]string
}

func (flags Flags) CheckIfFlagsAreSet(wanted ...string) bool {
	for _, wantedFlag := range wanted {
		if flags.Flag[wantedFlag] == "" {
			return false
		}
	}
	return true
}
