//1120009 - Feliciana Gunadi
package controllers

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Email    string `json:"e-mail"`
	UserType int    `json:"user type"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"e-mail"`
	UserType int    `json:"user type"`
}

type UserResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    UserInfo `json:"data"`
}

type UsersResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []UserInfo `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
