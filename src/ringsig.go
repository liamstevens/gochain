package gochain

import (
	"crypto/sha256"
)

type RingSignature struct {
	k []int
	l int
	n int
	q int
	p []byte
}

func New(k []int, L int) RingSignature {
	var q int
	if L-1 > 1 {
		q = L - 1
	} else {
		q = 1
	}
	r := RingSignature{k, L, len(k), q, []byte{}}
	return r
}

func permutate(ring RingSignature, m string) {
	h := sha256.New()
	h.Write([]byte(m))
	ring.p = h.Sum(nil)
}

func Verify(ring RingSignature, m string, X []byte) {
	permutate(ring, m)

}

func f(ring RingSignature, i int) {}
