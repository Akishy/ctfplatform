package entities

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (u *User) SetHashedPassword() error {
	bytesPassword := []byte(u.Password)
	hashedBytes, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	hash := string(hashedBytes[:])
	u.Password = hash
	return nil
}

func (u *User) ComparePasswords(otherPassword string) bool {
	incoming := []byte(otherPassword)
	existing := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err == nil
}
