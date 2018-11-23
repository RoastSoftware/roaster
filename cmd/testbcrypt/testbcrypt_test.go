package testbcrypt_test

import (
	"fmt"
	"testing"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/cmd/testbcrypt"
	"golang.org/x/crypto/bcrypt"
)

func benchmarkBcrypt(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		testbcrypt.HashWithSamePassword(i)
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

func TestExpectNoNullByteBug(t *testing.T) {
	res1 := testbcrypt.HashWithSameCost("abc\x00def")
	if err := bcrypt.CompareHashAndPassword(res1, []byte("abc\x00ghi")); err != nil {
		fmt.Println("Go does _not_ have the null byte bug 👍")
	} else {
		t.Error("Go has the null byte bug 👎")
	}
}

func TestExpectTruncatePassword(t *testing.T) {
	s72b := "104751087632048762130750394869807019873409852374095283740952837423452345"
	res1 := testbcrypt.HashWithSameCost(s72b + "A")
	if err := bcrypt.CompareHashAndPassword(res1, []byte(s72b+"B")); err != nil {
		// Bcrypt _should_ truncate passwords.
		t.Error("Go does _not_ truncate passwords after 72 bytes 👍")
	} else {
		// We expect the passwords to be truncated (part of the Blowfish
		// spec.)
		fmt.Println("Go truncates passwords after 72 bytes 👎")
	}
}
