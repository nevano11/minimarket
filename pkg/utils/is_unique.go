package utils

func IsUnique[T comparable](sl []T) bool {
	m := make(map[T]struct{}, len(sl))
	for _, item := range sl {
		if _, hasValue := m[item]; hasValue {
			return false
		}
		m[item] = struct{}{}
	}
	return true
}
