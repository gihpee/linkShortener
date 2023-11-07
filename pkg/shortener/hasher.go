package shortener

import (
	"crypto/md5"
	"encoding/hex"
)

func short(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash[:6] + "_.Go"
}
