package models

type User struct {
	Id        int                `json:"id" gorm:"primaryKey"`
	Name      string             `json:"name"`
	LastName  string             `json:"last_name"`
	UserName  string             `json:"user_name"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
	CreatedAt int                `json:"created_at"`
	UpdatedAt int                `json:"updated_at"`
	Purchases []Purchase         `json:"purchases" gorm:"foreignKey:UserID"`
	Category  []CategoryPurchase `json:"category" gorm:"foreignKey:UserID"`
}
