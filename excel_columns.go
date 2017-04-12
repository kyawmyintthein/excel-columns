package excelColumns

import "fmt"

func GetNextColumn(name string) string {
	slice := []rune(name)
	if len(slice) == 1 {
		runeNumber := slice[0]
		if runeNumber == Z {
			arr := []rune{A, A}
			return string(arr)
		} else {
			runeNumber = runeNumber + 1
			arr := []rune{runeNumber}
			return string(arr)
		}
	} else {
		runeNumber := slice[len(slice)-1]
		if runeNumber == Z {
			var lastIndex int
			for index, _ := range slice {
				if slice[index] == Z {
					lastIndex = index - 1
					break
				}
			}
			count := len(slice)
			if lastIndex == 0 {
				count = len(slice) - 1
			}
			for i := lastIndex; i <= count-lastIndex; i++ {
				if lastIndex == i {
					runeNumber := slice[lastIndex]
					runeNumber = runeNumber + 1
					slice[i] = runeNumber
				} else {
					slice[i] = A
				}
			}
		} else {
			runeNumber = runeNumber + 1
			slice = slice[:len(slice)-1]
			slice = append(slice, runeNumber)
			return string(slice)
		}
	}
	return ""
}
