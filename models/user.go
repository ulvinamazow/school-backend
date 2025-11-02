package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;not null"`
	Username string `json:"username" gorm:"not null:default: null"`
	Name     string `json:"name" gorm:"not null:default: null"`
	Surname  string `json:"surname" gorm:"not null:default: null"`
	Email    string `json:"email" gorm:"not null;default: null"`
	RoleID   int    `json:"roleid" gorm:"not null:default: null"`
}
