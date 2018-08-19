package funs

// InArr checks if an item is in a array
func InArr(testArr []string, toFind string) bool {
	for _, item := range testArr {
		if item == toFind {
			return true
		}
	}
	return false
}
