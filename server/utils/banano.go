package utils

import (
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/bbedward/nano/address"
	"github.com/bbedward/nano/types"
)

const bananoRegexStr = "(?:ban)(?:_)(?:1|3)(?:[13456789abcdefghijkmnopqrstuwxyz]{59})"

var bananoRegex = regexp.MustCompile(bananoRegexStr)

func GenerateAddress() string {
	pub, _ := address.GenerateKey()
	return strings.Replace(string(address.PubKeyToAddress(pub)), "nano_", "ban_", -1)
}

// ValidateAddress - Returns true if a nano address is valid
func ValidateAddress(account string) bool {
	if !bananoRegex.MatchString(account) {
		return false
	}
	return address.ValidateAddress(types.Account(account))
}

// Convert address to pubkey
func AddressToPub(account string) string {
	pubkey, _ := address.AddressToPub(types.Account(account))
	return hex.EncodeToString(pubkey)
}
