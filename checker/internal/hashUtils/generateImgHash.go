package hashUtils

import "golang.org/x/crypto/bcrypt"

func GenerateImgHash(codeArchive string) (string, error) {
	bytesImg := []byte(codeArchive)

	hashedBytes, err := bcrypt.GenerateFromPassword(bytesImg, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}
