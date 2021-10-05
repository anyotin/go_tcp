package Entities

type User struct {
	UserId   MuNumber  `json:"user-id"`
	UserName MyString `json:"user_name"`
	UserRank MuNumber  `json:"user_rank"`
}

type MuNumber int16

type MyString string
