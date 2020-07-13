package main

import (
	"fmt"

	"github.com/vickxxx/crypt/aes"
)

func main() {
	bkey := []byte("1111111111111111")
	var k [16]byte
	copy(k[:], bkey)
	txt := "1234"
	cipher, err := aes.Aes128ECBEncrypt(k, []byte(txt))
	fmt.Println(cipher, err)

	t, err := aes.Aes128ECBDecrypt(k, cipher)
	fmt.Println(string(t), err)
}
