package util

import (
	"math/rand"
	"time"
)

//随机截取字符串n长度
func RandomString(n int) string {
	var letters = []byte("asvdgahahsbdagaskdashasdknjsal")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
