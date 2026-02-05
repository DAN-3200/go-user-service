package ports

type IServices interface {
	SendMail(to string, body string) error
	GenerateUUID() string
	HashPassword(password string) (string, error)
	CompareHashPassword(pivotPassword, inputPassword string) error
	GenerateJWT(userID string, userRole string) (string, error)
}

type ISession interface {
	SetUserSession(Id string, Name string, Email string, Role string, JWT string) error
	LogoutUserSession(Id string) error
}
