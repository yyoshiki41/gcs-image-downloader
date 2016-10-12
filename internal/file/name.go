package file

import (
	"math/rand"
	"path"
	"path/filepath"
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

func Name(link string) string {
	name := randName()

	_, fileName := path.Split(link)
	ext := filepath.Ext(fileName)

	return name + ext
}

func randName() string {
	n := make([]byte, length)
	for i := range n {
		c := rand.Int63n(base)
		n[i] = strconv.FormatInt(c, base)[0]
	}
	return string(n)
}
