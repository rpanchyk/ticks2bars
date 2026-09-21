package utils

func FirstNonEmptyString(values ...string) string {
	for _, val := range values {
		if val != "" {
			return val
		}
	}
	return ""
}

func FirstNonZeroInt(values ...int) int {
	for _, val := range values {
		if val != 0 {
			return val
		}
	}
	return 0
}

func FirstNonFalseBool(values ...bool) bool {
	for _, val := range values {
		if val {
			return true
		}
	}
	return false
}
