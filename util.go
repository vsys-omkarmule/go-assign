package util

func add(a string) string {
	joinword = joinword + a + ","
	return joinword
}

// It takes arguments as string and seprator
func split(a, b string) string {
	for i := 0; i < len(a); i++ {

		if string(a[i]) == b {
			splitword = splitword + " "
		} else {
			splitword = splitword + string(a[i])
		}

	}

	return splitword
}

