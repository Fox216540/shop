package hasher

type Hasher interface {
	Hash(password string) (string, error)
	Verify(password string, hash string) (bool, error)
}
