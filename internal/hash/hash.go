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

func Verify(b []byte, sha1_hash string) bool {
	return sha1_hash == New(b)
}

func Read(sha1_hash string) []byte {
	b, err := hex.DecodeString(sha1_hash)
	if err != nil {
		panic(err)
	}
	return b
}
