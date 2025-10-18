package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil) // calling non args func
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
	/*
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
		case reflect.Map:
			for _, key := range val.MapKeys() {
				walk(val.MapIndex(key).Interface(), fn)
			}
		}

		for i := 0; i < numberOfValues; i++ {
			walk(getField(i).Interface(), fn)
		}

	*/
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
