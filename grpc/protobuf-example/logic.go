package gravatar

import (
	"crypto/md5"
	"fmt"
)

func gravatarHash(email string) []byte {
	return md5.Sum([]byte(email))[:]
}

func gravatarURL(hash [16]byte, size uint32) string {
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d", hash, size)
}

func gravatar(email string, size uint32) string {
	hash := gravatarHash(email)

	return gravatarURL(hash, size)
}
