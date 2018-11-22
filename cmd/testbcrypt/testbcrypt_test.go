package main

import "testing"

func benchmarkBcrypt(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		HashWithSamePassword(i)
	}
}

func BenchmarkBcrypt9(b *testing.B) {
	benchmarkBcrypt(9, b)
}

func BenchmarkBcrypt10(b *testing.B) {
	benchmarkBcrypt(10, b)
}

func BenchmarkBcrypt11(b *testing.B) {
	benchmarkBcrypt(11, b)
}

func BenchmarkBcrypt12(b *testing.B) {
	benchmarkBcrypt(12, b)
}

func BenchmarkBcrypt13(b *testing.B) {
	benchmarkBcrypt(13, b)
}

func BenchmarkBcrypt14(b *testing.B) {
	benchmarkBcrypt(14, b)
}
