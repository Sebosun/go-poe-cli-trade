package helpers

func SliceJoinStrings(s []string, separator string) string {
	acc := ""

	for i, string := range s {
		// last item in slice
		if len(s)-1 == i {
			acc += string
			continue
		}
		acc += string + separator
	}
	return acc
}
