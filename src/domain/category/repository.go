package category

type Repository interface {
	FindAll() ([]Category, error)
}
