package entity

type User struct {
	Id       string `db:"id"`
	Name     string `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type RegisterResponse struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	AccessToken string `json:"accesstoken"`
}
