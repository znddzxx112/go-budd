package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func Md5SumWithString(hash string, salt string) string {
	h := md5.New()
	h.Write([]byte(hash + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateTicket(pfx string, l int) string {
	rand.Seed(time.Now().UnixNano())
	var TicketRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	t := make([]rune, l)
	for i := range t {
		t[i] = TicketRunes[rand.Intn(len(TicketRunes))]
	}
	if pfx == "" {
		return fmt.Sprintf("%s", string(t))
	}
	return fmt.Sprintf("%s-%s", pfx, string(t))
}
