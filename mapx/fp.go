package mapx

import "maps"

/* Functional utilities for maps */

func ForEach[K comparable, V any](m map[K]V, action func(K, V)) {
	for k, v := range m {
		action(k, v)
	}
}
func Apply[K comparable, V any](m map[K]V, fn func(K, V) V) map[K]V {
	result := make(map[K]V, len(m))

	for k, v := range m {
		result[k] = fn(k, v)
	}
	return result
}
func Map[K, NK comparable, V, NV any](m map[K]V, mapper func(K, V) (NK, NV)) map[NK]NV {
	result := make(map[NK]NV, len(m))

	for k, v := range m {
		newKey, newVal := mapper(k, v)
		result[newKey] = newVal
	}
	return result
}
func Flatten[K comparable, V, S any](m map[K]V, flattener func(K, V) S) []S {
	result := make([]S, 0, len(m))

	for k, v := range m {
		result = append(result, flattener(k, v))
	}

	return result
}

type pair[T, U any] struct {
	First  T
	Second U
}

/*
	Default values behavior:
	 - defaultValues are consumed in order for each missing key.
	 - If there are more missing keys than default values, the last default value
	   is reused for all remaining keys.

Example:

	mainMap:   {a:1, b:2, c:3}
	zippedMap: {a:10}
	defaults:  [100, 200]

Result:

	a → (1,10)
	b → (2,100)
	c → (3,200)

If defaults were [100]:

	a → (1,10)
	b → (2,100)
	c → (3,100)
*/
func Zip[K comparable, V1, V2 any](mainMap map[K]V1, zippedMap map[K]V2, defaultValues ...V2) map[K]pair[V1, V2] {
	result := make(map[K]pair[V1, V2], len(mainMap))
	defaultIndex := 0

	for k, v1 := range mainMap {
		if v2, exists := zippedMap[k]; exists {
			result[k] = pair[V1, V2]{First: v1, Second: v2}
		} else if len(defaultValues) > 0 {
			result[k] = pair[V1, V2]{First: v1, Second: defaultValues[defaultIndex]}
			defaultIndex = min(defaultIndex+1, len(defaultValues)-1)
		} else {
			result[k] = pair[V1, V2]{First: v1}
		}
	}
	return result
}
func Extend[K comparable, V any](firstMap map[K]V, secondMap map[K]V) map[K]V {
	newMap := make(map[K]V, len(firstMap)+len(secondMap))
	maps.Copy(newMap, firstMap)

	for k, v := range secondMap {
		if _, exists := newMap[k]; exists {
			continue
		}

		newMap[k] = v
	}
	return newMap
}
func Reduce[K comparable, V, R any](m map[K]V, reducer func(R, K, V) R, initial R) R {
	result := initial

	for k, v := range m {
		result = reducer(result, k, v)
	}
	return result
}
