package stringz // import "ezpkg.io/stringz"

func FirstNParts(s string, n int, sep byte) string {
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			n--
			if n <= 0 {
				return s[i+1:]
			}
		}
	}
	return s
}

func LastNParts(s string, n int, sep byte) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == sep {
			n--
			if n <= 0 {
				return s[i+1:]
			}
		}
	}
	return s
}
