package idgenerator

type Generator interface {
	NewID() (string, error)
}
