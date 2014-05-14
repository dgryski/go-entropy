// Package entropy implements a streaming algorithm for entropy estimation
/*

http://jmlr.org/proceedings/papers/v31/clifford13a.pdf

It also implements an exact entropy calculator.

*/
package entropy

import (
	"math"
	"math/rand"

	"github.com/dgryski/go-spooky"
)

type RandFloat64 interface {
	Float64() float64
}

// Sketch is a sketch for estimating the entropy of a stream
type Sketch struct {
	Y int       // total count of elements
	y []float64 // elements of the sketch
}

// New returns a new sketch.  With k=O(1/eps**2), produces an epsilon-accurate sketch.
func NewEstimate(k int) *Sketch {
	return &Sketch{
		y: make([]float64, k),
	}
}

// Push adds element b to the stream dt times.
func (s *Sketch) Push(b []byte, dt int) {

	it := spooky.Hash64(b, 0)

	// Line 4
	s.Y += dt

	dtf := float64(dt)

	// Line 5
	rsrc := rand.New(rand.NewSource(int64(it)))

	// Line 6
	for j := range s.y {
		// Line 7
		rjit := maxSkew(rsrc)
		// Line 8
		s.y[j] += rjit * dtf
	}
}

// Entropy returns an estimate of the entropy of the stream so far
func (s *Sketch) Entropy() float64 {

	sum := 0.0
	for _, y := range s.y {
		// Line 9
		y /= float64(s.Y)
		// Line 10
		sum += math.Exp(y)
	}

	return -math.Log2(sum / float64(len(s.y)))
}

// maxSkew return a float from the maximally skewed stable distribution F(x;1,-1,math.Pi/2,0)
func maxSkew(r RandFloat64) float64 {

	// Table 1

	u1 := r.Float64()
	u2 := r.Float64()

	// This math, specifically the Log2 calls, is the algorithm's bottleneck

	w1 := math.Pi * (u1 - 0.5)
	w2 := -math.Log2(u2)

	halfPiW1 := math.Pi/2 - w1

	return math.Tan(w1)*(halfPiW1) + math.Log2(w2*(math.Cos(w1)/halfPiW1))
}
