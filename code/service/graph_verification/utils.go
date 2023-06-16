package graph_verification

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTicket(pfx string, l int) string {
	rand.Seed(time.Now().UnixNano())
	var TicketRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	t := make([]rune, l)
	for i := range t {
		t[i] = TicketRunes[rand.Intn(len(TicketRunes))]
	}

	return fmt.Sprintf("%s-%s", pfx, string(t))
}
