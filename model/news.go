package model

import "time"

type OneNews struct {
	ID			int			`json:"id" validate:"required"`
	Header		string		`json:"header" validate:"required"`
	Description	string		`json:"description" validate:"required"`
	URL 		string		`json:"url" validate:"required"`
	Site		string		`json:"site" validate:"required"`
	ImageSrc 	string		`json:"imagesrc" validate:"required"`
	Date		time.Time	`json:"date" validate:"required"`
}

