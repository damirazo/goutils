package helper

import "strconv"

// Приведение числа к строке
func int2str(value int) string {
	return string(strconv.Itoa(value))
}
