package db

type Model struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}

type Page struct {
	Total int64
	Data  interface{}
}
