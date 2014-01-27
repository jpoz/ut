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
