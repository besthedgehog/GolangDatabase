package main

// Спросить у нейросети, правильно ли я решил
// Функция проверсят, являются ли две строки зеркальными друг для друга
// func isAnagram(s string, t string) bool {
// 	lenOfWord := len(s)
// 	if lenOfWord != len(t) {
// 		return false
// 	}
// 	for i := range lenOfWord {
// 		if s[i] != t[lenOfWord-1-i] {
// 			fmt.Printf("i = %v, s[i] = %c, s[lenOfWord - 1 -i = %c\n", i, s[i], s[lenOfWord-1-i])
// 			return false
// 		}
// 	}
// 	return true
// }

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	storage1 := make(map[rune]int)
// 	storage2 := make(map[rune]int)
// 	for _, symb := range s {
// 		storage1[symb] += 1
// 	}
// 	for _, symb := range t {
// 		storage2[symb] += 1
// 	}
// 	for key, _ := range storage1 {
// 		if storage1[key] != storage2[key] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	counts := make(map[rune]int)
// 	for _, symb := range s {
// 		counts[symb]++
// 	}
// 	for _, symb := range t {
// 		counts[symb]--
// 		if counts[symb] < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// Массив на 26 элементов. Он создается прямо на стеке (очень быстро)
	var counts [26]rune
	for i := 0; i < len(s); i++ {
		// Трюк с ASCII: 'a' имеет код 97, 'b' - 98 и т.д.
		// Вычитая 'a', мы получаем индексы от 0 до 25!
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	// Проверим, что везде 0
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
