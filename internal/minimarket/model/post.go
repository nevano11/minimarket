package model

import (
	"fmt"
	"time"
)

type Post struct {
	Name        string    `json:"name"         validate:"min=3,max=40"`
	Description string    `json:"description"  validate:"min=3,max=100"`
	Price       int       `json:"price"        validate:"gt=0"`
	Image       *string   `json:"image"`
	CreatedAt   time.Time `json:"created-at"   db:"created_at"`
}

type PostAbout struct {
	Post
	Login         string `json:"login" db:"login"`
	IsCurrentUser bool   `json:"is_current_user" db:"is_current_user"`
}

func (x Post) String() string {
	img := "[empty image]"
	if x.Image != nil {
		img = *x.Image
	}

	return fmt.Sprintf("name=%s, image=%s, price=%d", x.Name, img, x.Price)
}
