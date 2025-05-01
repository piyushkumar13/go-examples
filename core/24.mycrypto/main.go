package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {

	secret := "myownsecret"

	file, err := os.ReadFile("sentinelsamplenovamicroservice2.sbom")

	if err != nil {
		log.Fatal("Not able to read", err)
	}

	data := file

	// create a new HMAC by defining the hash type and the key
	hmac := hmac.New(sha256.New, []byte(secret))

	// compute the HMAC
	hmac.Write([]byte(data))
	dataHmac := hmac.Sum(nil)

	hmacHex := hex.EncodeToString(dataHmac)
	secretHex := hex.EncodeToString([]byte(secret))

	fmt.Println("The hmac hash is ::: ", hmacHex)
	fmt.Println("The hmac secret is ::: ", secretHex)

	//fmt.Printf("HMAC_SHA256(key: %s, data: %s): %s", secretHex, hmacHex)
}
