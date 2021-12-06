package ringsig

type RingSignature struct {
	k []int
	l int
	n int
	q int
}

func New(k []int, L int) RingSignature {
	var q int
	if L-1 > 1 {
		q = L - 1
	} else {
		q = 1
	}
	r := RingSignature{k, L, len(k), q}
	return r
}

func permutate(ring RingSignature, m string) {

}
