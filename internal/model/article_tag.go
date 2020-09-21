package model

type ArticleTag struct {
	*Model
	ArticleID uint32 `json:"article_id"`
	TagID uint32 `json:"tag_id"`
}
