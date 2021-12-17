package file

import "gorm.io/gorm"

type File struct {
	gorm.Model
	URL string `json:"url"`
}
