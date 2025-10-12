package hasher

type Hasher interface {
	Hash(password string) (hash string, err error)
	Verify(password string, hash string) error
}
