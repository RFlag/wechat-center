package model

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	token = "wechat"
)

func CheckServer(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
