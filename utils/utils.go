package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const domain = "http://www.shawrtyUrl.in/"

func GenerateRandom(baseID int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomSuffix := make([]byte, 2)
	for i := range randomSuffix {
		randomSuffix[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Println("Random Suffix: ", string(randomSuffix))
	return EncodeBase62(baseID) + string(randomSuffix)
}

func EncodeBase62(num int) string {
	const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if num == 0 {
		return string(base62[0])
	}

	result := ""
	for num > 0 {
		result = string(base62[num%62]) + result
		num /= 62
	}
	return result
}

func GenerateURL(path string) string {
	return domain + path
}
