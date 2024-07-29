package models

type CategoryPurchase struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	UserID    int        `json:"user_id"`
	Compras   []Purchase `json:"compras" gorm:"foreignKey:CategoriaID"`
	CreatedAt int64      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64      `json:"updated_at" gorm:"autoUpdateTime"`
} //@Name CategoryPurchase

type PuncharseCategoryAndPurchases struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Compras   []Purchase `json:"compras" gorm:"foreignKey:CategoriaID"`
	CreatedAt int64      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64      `json:"updated_at" gorm:"autoUpdateTime"`
}

// control + j = command terminal
