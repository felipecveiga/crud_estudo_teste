
package model

type User struct {
	ID uint `json: "id" gorm:"primaryKey"`
	Nome string `json: "nome"`
	Idade int `json:"idade"`
	Email string `json:"email"`
}