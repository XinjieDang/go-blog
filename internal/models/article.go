package models

type Article struct {
	BasicModel
	Title        string `json:"title"`
	Content      string `json:"content"`
	Status       int64  `json:"status" gorm:"default:0"`
	AuthorId     int64  `json:"author_id"`
	CategoryId   int64  `json:"category_id"`
	ViewCount    int64  `json:"view_count"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	IsTop        int64  `json:"is_top" gorm:"default:0"`
	IsHot        int64  `json:"is_hot" gorm:"default:0"`
}
