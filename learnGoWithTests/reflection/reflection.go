package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field //Assign the val.Field method (function) to getField.
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
	/*
		if val.Kind() == reflect.Slice {
			for i := 0; i < val.Len(); i++ {
				walk(val.Index(i).Interface(), fn)
			}
			return
		}

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)

			switch field.Kind() {
			case reflect.String:
				fn(field.String())
			case reflect.Struct:
				walk(field.Interface(), fn)
			}
		}

	*/
	//field := val.Field(0) //use first and only field unsafe because what if there is not fields
	//fn(field.String())    // assumed the type is string, what if other than string
	//fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
