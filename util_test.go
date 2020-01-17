package jwt

import (
	"fmt"
	"testing"
)

func TestRandstr(t *testing.T) {
	var s = RandByte()
	fmt.Println(string(s))
}

