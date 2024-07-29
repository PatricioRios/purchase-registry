package models

// Article representa un artículo que pertenece a una compra
type Article struct {
	Id        int      `json:"id" gorm:"primaryKey"`
	Name      string   `json:"name"`
	Price     float64  `json:"price"`
	CompraID  uint     `json:"compra_id"`
	Compra    Purchase `json:"compra"`
	CreatedAt int64    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64    `json:"updated_at" gorm:"autoUpdateTime"`
}

type ArticleUpdate struct {
	Id        *int      `json:"id" gorm:"primaryKey"`
	Name      *string   `json:"name"`
	Price     *float64  `json:"price"`
	CompraID  *uint     `json:"compra_id"`
	Compra    *Purchase `json:"compra"`
	CreatedAt *int64    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *int64    `json:"updated_at" gorm:"autoUpdateTime"`
}
