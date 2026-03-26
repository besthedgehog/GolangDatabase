package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

const defaultCapacity uint64 = 1024

// node это струкура для отдельных нод
type node[K comparable, V any] struct {
	key     K
	value   V
	nextPtr *node[K, V]
}

// HashTable это структура для
type HashTable[K comparable, V any] struct {
	capacity uint64
	size     uint64
	table    []*node[K, V]
}

func (ht *HashTable[K, V]) newNode(key K, value V) *node[K, V] {
	return &node[K, V]{key: key,
		value:   value,
		nextPtr: nil}
}

func NewHashTable[K comparable, V any]() *HashTable[K, V] {
	return &HashTable[K, V]{
		capacity: defaultCapacity,
		size:     0,
		// Создаём слайс, инициализируем его дефолтными дначениями nil
		table: make([]*node[K, V], defaultCapacity),
	}
}

func NewHashTableWithCapacity[K comparable, V any](capacity uint64) (*HashTable[K, V], error) {
	if capacity == 0 {
		return nil, errors.New("capacity cannot be zero")
	}

	if capacity&(capacity-1) != 0 {
		return nil, errors.New("capacity should be a power of 2")
	}

	return &HashTable[K, V]{
		capacity: capacity,
		size:     0,
		table:    make([]*node[K, V], capacity),
	}, nil
}

func (hm *HashTable[K, V]) hash(key K) uint64 {

	// 1) Создаём объект-хешер с алгоритмом New64a
	// Он будет превращать любые объекты в число unit64
	h := fnv.New64a()

	// 2) Превращаем наш ключ типа K в строку, а строку в байты
	// Ошибки игнорируем, так как запись в память хешера всегда успешна
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	// 3) Получаем итоговое 64-битное число (хеш)
	hashValue := h.Sum64()

	// 4) Сдвигаем на 16 бит
	// Мы НЕ делаем хэш меньше (число всё еще 64-битное).
	// Мы просто создаем "копию" хэша, сдвинутую вправо.
	littleHash := hashValue >> 16

	// 5)  XOR (Перемешивание)
	// Мы смешиваем оригинал и копию.
	// Это нужно, чтобы "хвост" хэша (который мы скоро отрежем)
	// впитал в себя информацию из "головы" хэша.
	// Это делает распределение ключей по таблице более случайным.
	mixedHash := hashValue ^ littleHash

	// 6. Маскирование (втискивание в рамки массива).
	// Вместо медленного % (остаток от деления) используем быстрый & (битовое И).
	// Результат гарантированно попадет в диапазон [0...capacity-1].
	resultHash := mixedHash & (hm.capacity - 1)
	return resultHash
}

// Идея остатка от деления:
// Математически, остаток от деления любого числа на 1024 — это всегда число в диапазоне [0...1023].

// Если хэш = 1025, остаток = 1.
//
// Если хэш = 2048, остаток = 0.

// Если хэш = 500, остаток = 500.
//
// мы накладываем маску capacity - 1 на maskHash.
//
// maskHash — это данные (огромные).
//
// capacity - 1 — это трафарет (маленький).
//
// Мы прикладываем трафарет к огромному числу и видим только те биты, которые «пролезают» в дырки трафарета.
// Всё, что не влезло — отсекается. В итоге получается число, которое гарантированно меньше или равно 1023.

// Capacity возвразает Capacity
//
// Зачем нужны методы-геттеры (типа Capacity())?
// В Go (и в ООП вообще) это называется инкапсуляцией.
//
// Защита от дурака: Поле capacity в твоей структуре написано с маленькой буквы.
// Это значит, что оно private (недоступно из других пакетов).
// Если бы ты разрешил всем менять hm.capacity = 0 напрямую, твоя хэш-таблица мгновенно «взорвалась»
// бы при следующем расчете хэша.
//
// Контроль: Метод Capacity() позволяет читать значение, но не дает его менять.
//
// Гибкость: Представь, что через год ты решишь, что capacity не хранится в поле, а вычисляется динамически. Если у тебя есть метод, ты просто поменяешь код внутри него. Если все читали поле напрямую — тебе придется переписывать весь проект.
//
// Вывод: Это не просто «красиво», это архитектурно грамотно. Мы даем пользователю интерфейс,
// а не голые внутренности структуры.
func (hm *HashTable[K, V]) Capacity() uint64 {
	return hm.capacity
}

// resolvePutCollision
func (hm *HashTable[K, V]) resolvePutCollision(key K, value V, index uint64) {
	hm.table[index] = &node[K, V]{
		key:     key,
		value:   value,
		nextPtr: hm.table[index],
	}
	fmt.Printf("Collision: {index: %v, key: %v, value: %+v}\n", index, key, value)
}

// Put вставляет в таблицу значение п индексу
func (hm *HashTable[K, V]) Put(key K, value V) {
	index := hm.hash(key)

	// Если по индексу пусто
	if hm.table[index] == nil {
		hm.table[index] = hm.newNode(key, value)
		hm.size++
	} else {
		// Если индекс уже есть
		for it := hm.table[index]; it != nil; {
			// Замена на новое значение
			if it.key == key {
				it.value = value
				return
			}
			it = it.nextPtr
		}
		// Слчай коллизии
		hm.resolvePutCollision(key, value, index)
		hm.size++
	}
}

func (hm *HashTable[K, V]) Get(key K) (V, error) {
	index := hm.hash(key)

	node := hm.table[index]

	// Если значения по индексу нет
	if node == nil {
		// Здесь не nil из-за generic типа value
		return *new(V), errors.New("key not found")
	}
	for node != nil {
		if node.key == key {
			return node.value, nil
		}
		node = node.nextPtr
	}
	return *new(V), errors.New("key not found")
}

// Contains проверяет, есть ли элемент по ключу
func (hm *HashTable[K, V]) Contains(key K) bool {
	index := hm.hash(key)

	node := hm.table[index]

	if node == nil {
		return false
	}

	for node != nil {
		if node.key == key {
			return true
		}
		node = node.nextPtr
	}
	return false
}

// Remove удалят элемент по ключу
func (hm *HashTable[K, V]) Remove(key K) error {
	index := hm.hash(key)

	// Проверка на пустую ячейку
	if hm.table[index] == nil {
		return errors.New("key not found")
	}

	// Если удаляем самый первый элемент в цепочке
	if hm.table[index].key == key {
		hm.table[index] = hm.table[index].nextPtr
		hm.size--
		return nil
	}

	// Идём по остальному списку (паттерн два указателя)
	prev := hm.table[index]
	curr := prev.nextPtr

	for curr != nil {
		if curr.key == key {
			// Перекидываем мост через удаляемую ноду
			prev.nextPtr = curr.nextPtr
			hm.size--
			return nil
		}
		prev = curr
		curr = curr.nextPtr
	}
	return errors.New("key not found")
}
