package entropy

import (
	"testing"
)

func TestEntropy(t *testing.T) {

	stream := "hello, world this is my saaaaaaaaammple"

	ex := NewExact()
	sk := NewEstimate(10000)

	for _, s := range stream {
		ex.Push([]byte{byte(s)}, 1)
		sk.Push([]byte{byte(s)}, 1)
	}

	t.Log(ex.Entropy())
	t.Log(sk.Entropy())

}
