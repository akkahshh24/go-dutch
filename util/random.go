package util

import "github.com/brianvoe/gofakeit/v6"

func RandomString(n uint) string {
	return gofakeit.LetterN(n)
}

func RandomNumber(min, max int) int {
	return gofakeit.IntRange(min, max)
}
