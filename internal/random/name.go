package random

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	base   = 36
	length = 16
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Name() string {
	n := make([]byte, length)
	for i := range n {
		c := rand.Int63n(base)
		n[i] = strconv.FormatInt(c, base)[0]
	}
	return string(n)
}
