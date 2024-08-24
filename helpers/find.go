package helpers

func Find[E any](x []E, compFunc func(elem E) bool) (int, bool) {
	for i, v := range x {
		if compFunc(v) {
			return i, true
		}

	}
	return -1, false
}

func FindStr(x []string, target string) (int, bool) {
	for i, v := range x {
		if v == target {
			return i, true
		}

	}
	return 0, false
}
