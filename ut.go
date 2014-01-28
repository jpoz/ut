package ut

// Any takes a slice and a func. If the func ever returns true Any will
// return true
//
//      s := []int{1,2,3,4,5,6,7}
//      ut.Any(s, func(i interface{}) bool {
//          return i.(int) > 3
//      }) // true
//
func Any(s interface{}, f func(interface{}) bool) bool {
	for _, obj := range s.([]interface{}) {
		if f(obj) {
			return true
		}
	}
	return false
}

// Contains checks a slice to see if an object exists inside of it. Note the
// slice can be of any type
//
//      s := []string{"hello","world"}
//      ut.Contains(s, "hello") // true
//      s = []interface{}{"hello","world", 1000}
//      if ut.Contains(s, 1000) // true
//
func Contains(s interface{}, o interface{}) bool {
	for _, obj := range s.([]interface{}) {
		if obj == o {
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
