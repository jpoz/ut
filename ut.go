package ut

//
func Contains(s interface{}, o interface{}) bool {
	for _, obj := range s.([]interface{}) {
		if obj == o {
			return true
		}
	}
	return false
}

func Any(s interface{}, f func(interface{}) bool) bool {
	for _, obj := range s.([]interface{}) {
		if f(obj) {
			return true
		}
	}
	return false
}
