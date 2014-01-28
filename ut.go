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

func Pluck(s []map[string]interface{}, str string) []interface{} {
	var returned []interface{}

	for _, mp := range s {
		o := mp[str]
		if o != nil {
			returned = append(returned, o)
		}
	}

	return returned
}

func SlicesEqual(a interface{}, b interface{}) bool {
	if len(a.([]interface{})) != len(b.([]interface{})) {
		return false
	}
	for i, v := range a.([]interface{}) {
		if v != b.([]interface{})[i] {
			return false
		}
	}
	return true
}
