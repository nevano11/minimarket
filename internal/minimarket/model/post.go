package model

import "fmt"

type Post struct {
	Name  string  `json:"name"  validate:"min=3,max=40"`
	Price int     `json:"price" validate:"gt=0"`
	Image *string `json:"image"`
}

type PostAbout struct {
	Post
	IsCurrentUser bool `json:"is_current_user" db:"is_current_user"`
}

func (x Post) String() string {
	img := "[empty image]"
	if x.Image != nil {
		img = *x.Image
	}

	return fmt.Sprintf("name=%s, image=%s, price=%d", x.Name, img, x.Price)
}
