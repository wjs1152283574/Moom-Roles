package tool

func InSlice(src []string, target string) bool {
	if len(src) <= 0 {
		return false
	}

	for _, v := range src {
		if target == v {
			return true
		}
	}

	return false
}
