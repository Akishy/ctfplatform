package hashUtils

import "golang.org/x/crypto/bcrypt"

func CompareImgHash(rawCodeArchive string, hashedImg string) bool {
	incoming := []byte(rawCodeArchive)
	existing := []byte(hashedImg)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err == nil
}
