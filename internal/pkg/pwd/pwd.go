package pwd

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"
)

const charset = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func SignSalt() string {
	salt := make([]byte, 32)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range salt {
		randIdx, _ := rand.Int(rand.Reader, charsetLen)
		salt[i] = charset[randIdx.Int64()]
	}
	return string(salt)
}

func SignSaltPwd(pwd, salt string) string {
	hash := hmac.New(md5.New, []byte("hGZU8nBYE0YZcMTKMb9q"))
	hash.Write([]byte(pwd + salt))
	encrypted := hex.EncodeToString(hash.Sum([]byte("")))
	return strings.ToUpper(encrypted)
}
