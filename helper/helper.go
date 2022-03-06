package helper

import "math/rand"

// Random string generator of numbers
func RandomString(n int) string {
	var letters = []rune("0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
