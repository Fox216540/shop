package hasherinfra

import (
	"golang.org/x/crypto/bcrypt"
	hash "shop/src/domain/hasher"
)

type hasher struct {
	h hash.Hasher
}

func NewHasher(h hash.Hasher) hash.Hasher {
	return &hasher{h: h}
}

func (h *hasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (h *hasher) Verify(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
