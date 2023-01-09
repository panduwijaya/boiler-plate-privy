// Package entity
// Automatic generated
package entity

import (
	"time"
)

// Cake entity
type Cake struct{
	ID int	`db:"id,omitempty" json:"id"`
	Title string	`db:"title,omitempty" json:"title"`
	Description string	`db:"description,omitempty" json:"description"`
	Rating *int	`db:"rating,omitempty" json:"rating"`
	Image string	`db:"image,omitempty" json:"image"`
	CreatedAt *time.Time	`db:"created_at,omitempty" json:"created_at"`
	UpdatedAt *time.Time	`db:"updated_at,omitempty" json:"updated_at"`
}
