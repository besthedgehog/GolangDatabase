package main

import (
	"reflect"
	"slices"
	"testing"
)

// Тесты написаны для исправленной функции, просто
// возвращаем получившийся список
func TestRemoveDuplicates(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		intput := []int{1, 1, 2}
		expected := []int{1, 2, 2}
		output := removeDuplicates(intput)
		if !slices.Equal(output, expected) {
			t.Error()
		}
	})
	t.Run("Test 2", func(t *testing.T) {
		intput := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		output := removeDuplicates(intput)

		expected := []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}
		if !reflect.DeepEqual(output, expected) {
			t.Error()
		}
	})
}
