package model

type OneNews struct {
	ID			int			`json:"id" validate:"required"`
	Header		string		`json:"header" validate:"required"`
	URL 		string		`json:"url" validate:"required"`
	Site		string		`json:"site" validate:"required"`
	ImageSrc 	string		`json:"imagesrc" validate:"required"`
	Date		string		`json:"date" validate:"required"`
}

