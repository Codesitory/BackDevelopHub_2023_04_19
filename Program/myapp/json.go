package myapp

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HTTPHeaderStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
