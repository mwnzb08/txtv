package config

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func init () {
	fmt.Println("----------------> md5mix.go read successful")
}

func Md5MixEncryption (encryps string) string {
	encryptionAlgorithm :=md5.New()
	first := len(encryps)
	var finalNumber int
	for i := 2 ; i <= first ; i++  {
		if first%i == 0 {
			finalNumber = i
			break
		}
	}
	encryps = strconv.Itoa(first) + encryps + strconv.Itoa(finalNumber)
	encryptionAlgorithm.Write([]byte(encryps))
	return hex.EncodeToString(encryptionAlgorithm.Sum(nil))
}
