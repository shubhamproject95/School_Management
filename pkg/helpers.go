package pkg

import (
	"crypto/md5"
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"
)

func RandomString() string {
	now := time.Now().UTC().String()
	random := strconv.Itoa(rand.Int())

	// hash := md5.New().Sum([]byte(now + random))
	// return base64.StdEncoding.EncodeToString([]byte(hash))
	hash := md5.New()
	hash.Write([]byte(now + random))
	return base64.StdEncoding.EncodeToString([]byte(now + random))

}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
