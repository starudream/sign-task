package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

func MD5Hex(content string) string {
	return HashHex(md5.New(), content)
}

func SHA256Hex(content string) string {
	return HashHex(sha256.New(), content)
}

func HMAC256Hex(key, content string) string {
	return HashHex(hmac.New(sha256.New, []byte(key)), content)
}

func HashHex(h hash.Hash, content string) string {
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
