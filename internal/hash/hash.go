package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

func md5Bytes(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func sha1String(data []byte) string {
	hasher := sha1.New()
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func sha256String(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func sha512String(data []byte) string {
	hasher := sha512.New()
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func sha3String(data []byte) string {
	hasher := sha3.New384()
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func blake2bString(data []byte) string {
	hasher, err := blake2b.New256(nil)
	if err != nil {
		return ""
	}
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func blake2sString(data []byte) string {
	hasher, err := blake2s.New256(nil)
	if err != nil {
		return ""
	}
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
