package main

import (
	// we change the name of the imported package to save space
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	fmt.Println("data:", data)

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("sEnc:", sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("sDec:", string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("uEnc:", uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println("uDec:", string(uDec))
}
