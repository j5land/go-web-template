package entity

type Login struct {
	UserName	string  `json:"username" binding:"required"`
	Password	string  `json:"password" binding:"required"`
}

type User struct {
	Id			int	`json:id`
	UserName	string  `json:username`
	NickName	string  `json:nickname`
}