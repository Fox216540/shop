package hasher

type UseCase interface {
	HashPass(password string) (hash string, err error)
	VerifyPass(password string, hash string) error
}
