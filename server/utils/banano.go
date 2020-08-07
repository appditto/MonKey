package utils

import "github.com/bbedward/nano/address"

func GenerateAddress() string {
	pub, _ := address.GenerateKey()
	return string(address.PubKeyToAddress(pub))
}
