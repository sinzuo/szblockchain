package main

import (
	"bytes"
	"crypto/sha256"
	"math"
)

type PowerOfWork struct {
}

func (p *PowerOfWork) Run() int64 {
	var count int64
	for count = 1; count < math.MaxInt64; count++ {

		b := sha256.Sum256(bytes.Join([][]byte{[]byte("jiang"), HexToString(count)}, []byte{}))

		a := hashBijiao(b[:])
		if a == true {
			return count
		}
	}
	return 0
}
