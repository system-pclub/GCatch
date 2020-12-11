package genKill

import "reflect"

func CompareTwoMaps(map1 map[interface{}] interface{}, map2 map[interface{}] interface{}) bool {
	// reflect.DeepEqual will compare two values. Comparison of slice, maps or pointers will be handled recursively
	return reflect.DeepEqual(map1, map2)
}
