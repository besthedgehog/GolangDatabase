package main

import "fmt"

// Создаём новый тип на основе уже сущетвующего
type Set[T comparable] map[T]struct {
}

// NewSet is a factory
func NewSet[T comparable](vals ...T) *Set[T] {
	// make работает с новым типом,
	// потому что он основан на типе map[T]struct
	set := make(Set[T])

	for _, item := range vals {
		set.Add(item)
	}
	return &set
}

func (s *Set[T]) Add(value T) bool {
	oldLength := len(*s)
	// Добавляем в map по ключу value
	// пустую структуру
	(*s)[value] = struct{}{}
	return oldLength == len(*s)
}

func (s *Set[T]) Empty() bool {
	return len(*s) == 0
}

func (s *Set[T]) Size() int {
	return len(*s)
}

func (s *Set[T]) RemoveAll() {
	*s = *NewSet[T]()
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := (*s)[value]
	return ok
}

// Difference разность множеств
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	diff := NewSet[T]()
	for elem := range *s {
		if !other.Contains(elem) {
			diff.Add(elem)
		}
	}
	return diff
}

// Симметричая разность
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	symDiff := NewSet[T]()
	for elem := range *s {
		if !other.Contains(elem) {
			symDiff.Add(elem)
		}
	}
	for elem := range *other {
		if !s.Contains(elem) {
			symDiff.Add(elem)
		}
	}
	return symDiff
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	intersectionSet := NewSet[T]()
	if s.Size() < other.Size() {
		for elem := range *s {
			if other.Contains(elem) {
				intersectionSet.Add(elem)
			}
		}
	} else {
		for elem := range *other {
			if s.Contains(elem) {
				intersectionSet.Add(elem)
			}
		}
	}
	return intersectionSet
}

// Объединение
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := NewSet[T]()
	for elem := range *s {
		unionSet.Add(elem)
	}
	for elem := range *other {
		unionSet.Add(elem)
	}
	return unionSet
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.Size() < other.Size() {
		return false
	}
	for elem := range *s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s *Set[T]) ForEach(fn func(value T)) {
	if s.Size() == 0 {
		return
	}
	for key, _ := range *s {
		fn(key)
	}
}

func (s *Set[T]) PrintSet() {
	fmt.Print("Set: [")
	s.ForEach(func(value T) {
		fmt.Printf("%+v, ", value)
	})
	fmt.Println("]")
}

func main() {
	setA := NewSet(1, 2, 3, 4, 5, 6, 3, 3, 3, 2, 1, 4, 7, 8, 5)
	fmt.Println("----- setA ------")
	setA.PrintSet()
	setB := NewSet(10, 22, 1, 4, 21, 4, 5, 21, 11, 10)
	fmt.Println("----- setB ------")
	setB.PrintSet()
	fmt.Println("----- Contains ------")
	value := 10
	fmt.Printf("SetA contains %v == %v\n",
		value, setA.Contains(value))
	fmt.Printf("SetB contains %v == %v\n", value, setB.Contains(value))

	fmt.Println("----- A - B ------")
	setA.Difference(setB).PrintSet()

	fmt.Println("----- B - A ------")
	setB.Difference(setA).PrintSet()

	fmt.Println("----- Union ------")
	setA.Union(setB).PrintSet()

	fmt.Println("----- Intersect ------")
	setA.Intersect(setB).PrintSet()
}
