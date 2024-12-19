package hashUtils

import "golang.org/x/crypto/bcrypt"

func CompareImgHash(codeArchive string, hashedImg string) bool {
	incoming := []byte(codeArchive)
	existing := []byte(hashedImg)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err == nil
}
