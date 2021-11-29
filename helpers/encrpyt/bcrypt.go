package encrpyt

import "golang.org/x/crypto/bcrypt"

func Hash(passwd string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
