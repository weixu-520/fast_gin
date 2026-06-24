package models

type UserModel struct {
	Model
	Username string `gorm:"size:16" json:"username"`
	Nickname string `gorm:"size:32" json:"nickname"`
	Password string `gorm:"size:64" json:"-"`
	RoleID   int8   `json:"roleID"`
}
