package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {

	key := "mykey"
	domain := "play.golang.org"

	secret := []byte(domain + "/" + key)

	for i := 0; i < 1<<16; i++ {
		sum := sha256.Sum256(secret)
		secret = sum[:]
	}

	password := base64.StdEncoding.EncodeToString(secret)[:16]
	fmt.Println(password)
	// fmt.Println("OIs/gtSGGsuWXLDJ") // expected

}
