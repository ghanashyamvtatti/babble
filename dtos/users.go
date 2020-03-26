package dtos

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegistration struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"name" binding:"required"`
}