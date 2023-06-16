package utils

import "testing"

func TestInetAtoN(t *testing.T) {
	t.Log(InetAtoN("218.109.200.119"))
	t.Log(InetNtoA(3664627831))
}
