package helpers

func FindStr(x []string, target string) (int, bool) {
	for i, v := range x {
		if v == target {
			return i, true
		}

	}
	return 0, false
}
