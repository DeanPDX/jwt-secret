// This is a utlity to create a base64 encoded HMAC signing
// secret that can then be used in env vars and such.
//
// From the `golang-jwt/jwt` [FAQ]:
//
//	If you have trouble handling a []byte key in our setup, e.g., because you
//	are reading it from your environment variables on your cluster or similar,
//	you can always use base64 encoding to have the key as a "string" type
//	outside of your program and then use base64.Encoding.DecodeString to
//	decode the base64 string into the []byte slice that the signing method needs.
//
// See also: [golang-jwt/jwt].
//
// [FAQ]: https://golang-jwt.github.io/jwt/usage/signing_methods/#frequently-asked-questions
// [golang-jwt/jwt]: https://github.com/golang-jwt/jwt
package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	// Print key options
	fmt.Print(`Which signing method are you using?
1. HS256
2. HS384
3. HS512
:`)
	// Read user input and set key length in bytes
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keyLength := 0
	switch scanner.Text() {
	case "1":
		keyLength = 32
	case "2":
		keyLength = 48
	case "3":
		keyLength = 64
	default:
		return
	}
	// Create slice based on key length and read random bytes into it
	key := make([]byte, keyLength)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Base64 encode and write to console
	encoded := base64.StdEncoding.EncodeToString(key)
	fmt.Printf("Base64 encoded:\n%v", encoded)
	// fmt.Println(encoded)
}
