package helper

import "golang.org/x/crypto/bcrypt"

func NormalizePassword(p string) []byte {
	return []byte(p)
}

func GenPassHash(p string) (str string, err error) {
	// Normalize password from string to []byte.
	bytePwd := NormalizePassword(p)

	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPwd, inputPwd string) bool {
	byteHash := NormalizePassword(hashedPwd)
	byteInput := NormalizePassword(inputPwd)

	// Return result.
	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
