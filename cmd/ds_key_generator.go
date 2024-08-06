package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"log"
	"os"
)

// Creates Public & Private Keys for Digital Signature inside deploy/keys/ directory
func main() {
	pub, prv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalln("error creating digital signature certificates")
	}

	_, err = os.OpenFile("./deploy/keys/public.key", os.O_RDONLY, 644)
	if err != nil {
		_ = os.WriteFile("./deploy/keys/public.key", pub, 0644)
		_ = os.WriteFile("./deploy/keys/private.key", prv, 0644)
	}
}
