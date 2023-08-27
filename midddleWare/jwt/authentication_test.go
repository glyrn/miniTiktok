package jwt

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	s := CreateToken(-1)
	fmt.Println(1)
	fmt.Println(s)
}
