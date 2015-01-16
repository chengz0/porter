package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash/fnv"
)

func FnvHash(data []byte) uint32 {
	hash := fnv.New32a()
	hash.Write(data)
	return hash.Sum32()
}

func SHA1Hash(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	ret := hash.Sum(nil)
	return fmt.Sprintf("%x", ret)
}

func MD5Hash(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	ret := hash.Sum(nil)
	return fmt.Sprintf("%x", ret)
}
