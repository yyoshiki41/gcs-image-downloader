package file

import (
	"math/rand"
	"net/url"
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

func Name(link string) (string, error) {
	name := randName()

	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	_, fileName := path.Split(u.Path)
	ext := filepath.Ext(fileName)

	return name + ext, nil
}

func randName() string {
	n := make([]byte, length)
	for i := range n {
		c := rand.Int63n(base)
		n[i] = strconv.FormatInt(c, base)[0]
	}
	return string(n)
}
