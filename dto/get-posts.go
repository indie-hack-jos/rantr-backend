package dto

type GetPosts struct {
    PageNumber string `json:"page_number" binding:"required"`
	PageSize string `json:"page_size" binding:"required"`
}