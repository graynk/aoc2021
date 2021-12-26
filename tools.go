package advent_of_code

import "unicode"

func ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(value int) int {
	if value < 0 {
		value *= -1
	}
	return value
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
