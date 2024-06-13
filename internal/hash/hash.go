package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

func New(b []byte) string {
	h := sha1.New()
	h.Write(b)
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
