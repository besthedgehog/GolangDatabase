package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	someSet := NewSet[int]()
	if someSet.Size() != 0 {
		t.Errorf("Length of the new set shoud be equal to 0")
	}
	someSet.Add(1)
	if someSet.Size() != 1 {
		t.Errorf("size should be 1, but %v", someSet.Size())
	}
	if !someSet.Contains(1) {
		t.Errorf("Error in contains func")
	}
	if someSet.Contains(8) {
		t.Errorf("8 is not in the set")
	}
}

func TestDifference(t *testing.T) {
	setA := NewSet[int](1, 2, 3)
	setB := NewSet[int](2, 3)
	resultSet := setA.Difference(setB)
	if resultSet.Size() != 1 {
		t.Errorf("Difference error")
	}
}

func TestSymmetricDifference(t *testing.T) {
	setA := NewSet[int](1, 2, 5)
	setB := NewSet[int](1, 2, 7)
	{
		resultSet := setA.SymmetricDifference(setB)
		if resultSet.Size() != 2 {
			t.Errorf("size of result set should be 2")
		}
		if !resultSet.Contains(5) {
			t.Errorf("Resutset should contain ")
		}
		if !resultSet.Contains(7) {
			t.Errorf("Resutset should contain ")
		}
	}
	{
		resultSet := setA.SymmetricDifference(setA)
		if resultSet.Size() != 0 {
			t.Errorf("Size should be 0")
		}
	}

}

func TestContains(t *testing.T) {
	setA := NewSet[string]("A", "B", "C")
	if !setA.Contains("A") {
		t.Errorf("Contains must return true")
	}
	if setA.Contains("E") {
		t.Errorf("Contains must return false")
	}

	setA.RemoveAll()
	if setA.Contains("A") {
		t.Errorf("Contains must return false")
	}
	if setA.Size() != 0 {
		t.Errorf("Error in RemoveAll func")
	}
}

func TestIntesect(t *testing.T) {
	setA := NewSet[int](1, 2, 3)
	setB := NewSet[int](3, 4, 5)
	resultSet := setA.Intersect(setB)
	if resultSet.Size() != 1 {
		t.Errorf("Size should be equal to 1")
	}
	if !resultSet.Contains(3) {
		t.Errorf("Should contain 3")
	}
}
