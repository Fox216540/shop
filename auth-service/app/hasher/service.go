package hasher

import domainHasher "github.com/Fox216540/shop/auth-service/domain/hasher"

type service struct {
	hasher domainHasher.Hasher
}

func NewService(h domainHasher.Hasher) UseCase {
	return &service{
		hasher: h,
	}
}

func (s *service) HashPass(password string) (string, error) {
	hash, err := s.hasher.Hash(password)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (s *service) VerifyPass(password string, hash string) error {
	if err := s.hasher.Verify(password, hash); err != nil {
		return err
	}
	return nil
}
