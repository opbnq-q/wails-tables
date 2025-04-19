package utils

import "reflect"

// ReplaceEmptySlicesWithNil takes a pointer to any struct (or struct-like type)
// and replaces all empty slices with nil, traversing nested fields recursively.
func ReplaceEmptySlicesWithNil(v interface{}) {
	// Must be a pointer so we can set fields
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic("ReplaceEmptySlicesWithNil expects a pointer to a struct")
	}

	replaceEmptySlices(rv.Elem())
}

// replaceEmptySlices is the recursive helper that does the actual reflection work.
func replaceEmptySlices(v reflect.Value) {
	switch v.Kind() {
	case reflect.Ptr:
		// If it is a pointer, recurse on the element it points to.
		if !v.IsNil() {
			replaceEmptySlices(v.Elem())
		}

	case reflect.Struct:
		// For a struct, iterate each field and process it.
		for i := 0; i < v.NumField(); i++ {
			fieldVal := v.Field(i)
			if !fieldVal.CanSet() {
				// If we can't set the field (e.g. unexported), skip it
				continue
			}
			replaceEmptySlices(fieldVal)
		}

	case reflect.Slice:
		// If it is a slice, check if it's empty; if it is, set it to nil.
		if v.Len() == 0 && v.CanSet() {
			// v.Set(reflect.Zero(v.Type())) sets slice to nil
			v.Set(reflect.Zero(v.Type()))
		} else {
			// If not empty, we should recurse on each element in case it's a struct.
			for i := 0; i < v.Len(); i++ {
				replaceEmptySlices(v.Index(i))
			}
		}
	}
}
