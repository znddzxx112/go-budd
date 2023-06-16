package utils

import (
	"testing"
)

func TestCheckPhone(t *testing.T) {
	phones := []string{
		"18868801234",
		"1886880123",
		"188688012344",
	}
	for _, p := range phones {
		pr := CheckPhone(p)
		t.Log(p, ":", pr)
	}
}
