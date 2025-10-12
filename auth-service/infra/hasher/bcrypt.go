package hasher

import (
	domainHasher "github.com/Fox216540/shop/auth-service/domain/hasher"
	"golang.org/x/crypto/bcrypt"
)

type hasher struct {
}

func NewHasher() domainHasher.Hasher {
	return &hasher{}
}

func (h *hasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", NewInvalidHashError(err)
	}

	return string(hashedPassword), nil
}

func (h *hasher) Verify(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return domainHasher.NewBadPasswordError(err)
	}
	return nil
}
