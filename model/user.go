
package model

type User struct {
	ID uint `json: "id" gorm:"primaryKey"`
	Nome string `json: "nome"`
	Idade int `json:"idade"`
	email string `json:"email"`
}