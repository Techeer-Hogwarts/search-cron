package models

type BaicIndex struct {
	ID    string `json:"id" form:"id" example:"1"`
	Index string `json:"index" form:"index" example:"index"`
}
