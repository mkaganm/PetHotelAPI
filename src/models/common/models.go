package common

type User struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	UserPassword string `json:"user_password"`
	UserEmail    string `json:"user_email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
}

func (u User) CheckRegisterRequest() bool {

	condition := u.Username == "" ||
		u.UserPassword == "" ||
		u.UserEmail == "" ||
		u.FirstName == "" ||
		u.LastName == "" ||
		u.PhoneNumber == ""

	if condition {
		return false
	} else {
		return true
	}
}
