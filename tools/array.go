package tools

func InArray[T comparable](target T, arr []T) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if target == arr[i] {
			return true
		}
	}
	return false
}
