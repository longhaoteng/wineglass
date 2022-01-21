package api

type Pageable struct {
	Page int `json:"page" form:"page" binding:"required,gte=1"`
	Size int `json:"size" form:"size" binding:"required,gte=1,lte=50"`
}

type PageOptional struct {
	Page int `json:"page" form:"page" binding:"omitempty,gte=1"`
	Size int `json:"size" form:"size" binding:"omitempty,gte=1,lte=50"`
}
