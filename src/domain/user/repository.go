package user

type Repository interface {
	Add(user User) (User, error)
	Delete(user User) error
	GetByID(id string) (User, error)
	Update(user User) (User, error)
}
