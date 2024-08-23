package concat

import "strings"

// Concat takes an input string, seperator and split char and returns the concatonated string
func Concat(str string, sep string, split string) string {

	s := strings.Split(str, split)
	return strings.Join(s, sep)
}
