package entropy

import "math"

type Exact map[string]float64

func NewExact() Exact {
	return make(Exact)
}

func (e Exact) Push(b []byte, dt int) {
	e[string(b)] += float64(dt)
}

func (e Exact) Entropy() float64 {
	hm := 0.
	size := 0.0
	for _, c := range e {
		size += c
		hm += c * math.Log2(c)
	}

	return math.Log2(size) - hm/size
}
