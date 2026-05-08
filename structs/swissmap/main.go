package main

import (
	"fmt"
	"hash/maphash"
)

// groupSize В новой мапе данные хранятся не по одному, а "группами" по 8 штук.
// Восемь элементов типа uint8 занимают ровно 64 бита (8 байт).
// Современные 64-битные процессоры могут за один
// такт процессора проверить сразу все 8 ячеек
//
// empty
const (
	groupSize = 8
	empty     = 0x80 // 10000000 или 0b10000000
	deleted   = 0xFE // 11111110 или 0b11111110
)

type Group[K comparable, V any] struct {
	ctrl [groupSize]uint8 // Короткие части хеша для поиска и ускорения сравнения
	keys [groupSize]K     // Почему массивы, а не связанные списки? Для бОльше скорости
	vals [groupSize]V
}

type SwissMap[K comparable, V any] struct {
	groups []Group[K, V]
	seed   maphash.Seed // соль для хеша
}

func NewSwissMap[K comparable, V any](numGroups int) *SwissMap[K, V] {
	m := &SwissMap[K, V]{
		groups: make([]Group[K, V], numGroups),
		seed:   maphash.MakeSeed(),
	}
	for i := range m.groups {
		for j := range m.groups[i].ctrl {
			m.groups[i].ctrl[j] = empty
		}
	}
	return m
}

func (m *SwissMap[K, V]) hash(key K) (uint64, uint8) {
	// Инициализируем одъект для работы с хеш-функциями
	var h maphash.Hash

	// Задаём использованную соль
	h.SetSeed(m.seed)

	// Записываем наш ключ в буффер h
	h.WriteString(fmt.Sprint(key))

	// Считаем хеш буфера
	hash64 := h.Sum64()

	// Отрежем последние 7 бит
	h1 := hash64 >> 7

	// Хэш (hash64): 10101010 11011011 (какой-то случайный хэш)
	// Маска (0x7F): 00000000 01111111 (наша маска, слева куча нулей)
	// --------------------------------
	// Результат (&): 00000000 01011011
	h2 := uint8(hash64 & 0x7F)
	return h1, h2
}

func (m *SwissMap[K, V]) Put(key K, val V) {
	// Получим хеш
	h1, h2 := m.hash(key)

	// Получим индекс по хешу индекс группы
	groupIdx := h1 & (uint64(len(m.groups)) - 1)

	// Зачем мы это задаём?
	var targetGroup *Group[K, V]
	var targetIdx int = -1

	for {
		g := &m.groups[groupIdx]
		for i, ctrlByte := range g.ctrl {
			// Если ключ уже есть, то обновляем значение
			if ctrlByte == h2 && g.keys[i] == key {
				g.vals[i] = val
				return
			}
			// Есть пустое место и до этого не встречали tombstone
			if ctrlByte == empty {
				if targetIdx == -1 {
					targetGroup = g
					targetIdx = i
				}
				// Вставка данных
				targetGroup.ctrl[targetIdx] = h2
				targetGroup.keys[targetIdx] = key
				targetGroup.vals[targetIdx] = val
				return
			}
			// Запоминаем первый tombstone, но продолжаем искать
			// вдруг наш ключ лежит дальше по цепочке
			if ctrlByte == deleted && targetIdx == -1 {
				targetGroup = g
				targetIdx = i
			}
		}
		// Переходим в другую группу
		// Если мы в группе 5: (5 + 1) % 8 = 6. (Перешли в 6-ю).
		// Если мы в группе 7: (7 + 1) % 8 = 0.
		// (Магия! Мы плавно телепортировались в самое начало парковки, в нулевую группу!).
		groupIdx = (groupIdx + 1) % uint64(len(m.groups)-1)
	}
}

// Get возвращает значение по ключу
func (m *SwissMap[K, V]) Get(key K) (V, bool) {
	// Получим хеши по ключу
	h1, h2 := m.hash(key)
	groupIdx := h1 & (uint64(len(m.groups)) - 1)

	for {
		g := m.groups[groupIdx]
		for i, ctrlByte := range g.ctrl {
			if ctrlByte == h2 && g.keys[i] == key {
				return g.vals[i], true
			}
			// Если встретили empty, дальше искать смысла нет
			if ctrlByte == empty {
				var zero V
				return zero, false
			}
		}
		// Если в этой группе нет, ищем в следующей
		groupIdx = (groupIdx + 1) & (uint64(len(m.groups)) - 1)
	}
}

func (m *SwissMap[K, V]) Delete(key K) {
	h1, h2 := m.hash(key)
	groupIdx := h1 & (uint64(len(m.groups)) - 1)

	for {
		g := &m.groups[groupIdx]
		for i, ctrlByte := range g.ctrl {
			if ctrlByte == h2 && g.keys[i] == key {
				g.ctrl[i] = deleted
				// Очищаем ссылки
				g.keys[i] = *new(K)
				g.vals[i] = *new(V)
				return
			}
			if ctrlByte == empty {
				return
			}
		}
		groupIdx = (groupIdx + 1) & (uint64(len(m.groups)) - 1)
	}
}
