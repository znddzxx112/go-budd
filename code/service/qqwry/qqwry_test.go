package qqwry

import "testing"

func TestQQwry_Find(t *testing.T) {
	wry, err := NewQQwry("./resources/ip/qqwry.dat")
	if err != nil {
		t.Fatal(err)
	}
	defer wry.Close()
	findWry, err := wry.Find("218.109.200.119")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(findWry.Ip, findWry.City, findWry.Country)
}
