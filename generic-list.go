package utils

func Contains[K comparable](item K, list []K) bool {
	if list != nil && len(list) > 0 {
		for _, i := range list {
			if item == i {
				return true
			}
		}
	}
	return false
}
