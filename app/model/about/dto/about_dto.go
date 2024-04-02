package dto

type AboutDTO struct {
	Title       string `json:"title" binding:"required"`
	Subtitle    string `json:"subtitle" binding:"required"`
	Description string `json:"description" binding:"required"`
	Note        string `json:"note" binding:"required"`
	Image       string `json:"image" binding:"required"`
}
