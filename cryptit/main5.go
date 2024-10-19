package main

import (
	"fmt"

	"github.com/Priyanka-yadavv/cryptit/decrypt"
	"github.com/Priyanka-yadavv/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
