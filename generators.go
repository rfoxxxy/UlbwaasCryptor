package main

import "math/rand"

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func generateString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return encryptStringToHEX(string(b))
}
