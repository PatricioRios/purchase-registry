package models

// wagger gin model
//
//	@Description	Compra asdasd
type Purchase struct {
	Id int `json:"id" gorm:"primaryKey"`
	// @Description Titulo de la compra
	Title string `json:"title"`
	// @Description Descripcion de la compra
	Description string `json:"description"`
	// @Description
	Import      float64          `json:"import"`
	UserID      int              `json:"user_id"`
	CategoriaID int              `json:"categoria_id"` // Foreign key
	Category    CategoryPurchase `json:"category" gorm:"foreignKey:CategoriaID"`
	Articulos   []Article        `json:"articulos" gorm:"foreignKey:CompraID"`
	CreatedAt   int64            `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64            `json:"updated_at" gorm:"autoUpdateTime"`
} //	@Name	Purchase

type CompraUpdate struct {
	Id          *int              `json:"id"`
	Title       *string           `json:"title"`
	Description *string           `json:"description"`
	Import      *float32          `json:"import"`
	CategoriaID *int              `json:"categoria_id"` // Foreign key
	Category    *CategoryPurchase `json:"category" gorm:"foreignKey:CategoriaID"`
	Articulos   *[]Article        `json:"articulos" gorm:"foreignKey:CompraID"`
}
