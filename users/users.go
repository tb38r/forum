package users

type User struct {
	UserID          int
	Email           string
	Username        string
	Password        string
	Usertype        string
	ExternalLoginId string
}

type AuthUser struct {
	Email        string
	PasswordHash string
}
