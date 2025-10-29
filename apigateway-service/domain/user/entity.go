package user

type User struct {
	Email    string // User's email address
	Name     string // User's full name
	Phone    string // User's phone number
	Password string // User's password (should be hashed in practice)
	Address  string // User's address
}
