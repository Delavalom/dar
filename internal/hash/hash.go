package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

func New(b []byte) string {
	h := sha1.New()
	if _, err := h.Write(b); err != nil {
		panic(err)
	}
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func Verify(b []byte, sha1_hash string) bool {
	return sha1_hash == New(b)
}
