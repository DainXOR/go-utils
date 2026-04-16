package collections

import (
	"slices"
)

/* Functional utilities for slices */

// Removes elements that predicate returns false
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))

	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}
func Map[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, 0, len(slice))

	for _, value := range slice {
		result = append(result, mapper(value))
	}
	return result
}

// MapE is a variant of Map that drops the elements for which the mapper returns an error.
// First option: Skip errors, keeps mapping even if any value returns an error. Returns the last error.
func MapE[T, U any](slice []T, mapper func(T) (U, error), options ...bool) ([]U, error) {
	result := make([]U, 0, len(slice))
	skipErrs := false

	if len(options) > 0 {
		skipErrs = options[0]
	}

	var lastError error
	lastError = nil
	for _, value := range slice {
		mappedValue, err := mapper(value)

		if err != nil {
			if skipErrs {
				lastError = err
				continue
			}

			return result, err
		} else {
			result = append(result, mappedValue)
		}
	}
	return result, lastError
}
func ForEach[T any](slice []T, action func(int, T)) {
	for i, value := range slice {
		action(i, value)
	}
}
func Reduce[T, U any](slice []T, reducer func(U, T) U, initial U) U {
	result := initial

	for _, value := range slice {
		result = reducer(result, value)
	}
	return result
}
func ReduceE[T, U any](slice []T, reducer func(U, T) (U, error), initial U) (U, error) {
	result := initial

	for _, value := range slice {
		if result, err := reducer(result, value); err != nil {
			return result, err
		}
	}
	return result, nil
}
func Any[T any](slice []T, predicate func(T) bool) bool {
	return slices.ContainsFunc(slice, predicate)
}
func All[T any](slice []T, predicate func(T) bool) bool {
	cmp := true
	for _, v := range slice {
		cmp = cmp && predicate(v)
	}

	return cmp
}
