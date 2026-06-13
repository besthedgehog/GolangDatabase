package main

// import "fmt"

// func strStr(haystack string, needle string) int {
// 	lenH := len(haystack)
// 	lenN := len(needle)

// 	if lenH*lenN == 0 {
// 		return 0
// 	}

// 	// Бежим до тех пор, пока в haystack не осатется места для needle
// 	for i := 0; i <= lenH-lenN; i++ {
// 		currentSlice := haystack[i : i+lenN]
// 		fmt.Printf("needle = %v, currentSlice = %v, i=%v\n", needle, currentSlice, i)
// 		if currentSlice == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }

///
//
//
//
//
//
//
//
//
//

// func strStr(haystack string, needle string) int {
// 	lenH := len(haystack)
// 	lenN := len(needle)
// 	if lenH == 0 {
// 		return -1
// 	}
// 	if lenN == 0 {
// 		return 0
// 	}
// 	for hayIdx := 0; hayIdx <= lenH-lenN; hayIdx++ {
// 		if needle == haystack[hayIdx:hayIdx+lenN] {
// 			return hayIdx
// 		}
// 	}
// 	return -1
// }

func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	if hLen == 0 {
		return -1
	}
	if nLen == 0 {
		return 0
	}
	for hayIdx := 0; hayIdx <= hLen-nLen; hayIdx++ {
		needleIdx := 0
		for needleIdx < nLen && haystack[needleIdx+hayIdx] == needle[needleIdx] {
			needleIdx++
		}
		if needleIdx == nLen {
			return hayIdx
		}
	}
	return -1
}
