package utils

import (
	"testing"
)

func TestPassword(t *testing.T) {
	pass := Password("123456", "APPihWUl9uD6W4kI")
	t.Log(pass)
}

func TestCheckPassword(t *testing.T) {
	pass := []string{
		"Az123456",
		"aa123456",
		"Aa12345",
	}
	for _, p := range pass {
		pr := CheckPassword(p)
		t.Log(p, ":", pr)
	}
}
