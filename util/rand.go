package util

import (
	"math/rand/v2"
)

const dictAlphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandAlphanum(n int) string {
	return RandDict(dictAlphanum, n)
}

const dictHex = "ABCDEF0123456789"

func RandHex(n int) string {
	return RandDict(dictHex, n)
}

func RandDict(dict string, n int) string {
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = dict[rand.IntN(len(dict))]
	}
	return string(bs)
}
