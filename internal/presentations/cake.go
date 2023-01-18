// Package presentations
// Automatic generated
package presentations

type (
	// CakeQuery parameter
	CakeQuery struct {
		ID          int    `db:"id,omitempty" json:"id" url:"id,omitempty"`
		Title       string `db:"title,omitempty" json:"title" url:"title,omitempty"`
		Rating      *int   `db:"rating,omitempty" json:"rating"`
		Description string `db:"description,omitempty" json:"description" url:"description,omitempty"`
		Image       string `db:"image,omitempty" json:"image" url:"image,omitempty"`
		Paging
		PeriodRange
	}

	CakeID struct {
		ID int `db:"id,omitempty" json:"id" url:"id,omitempty"`
	}

	// CakeParam input param
	CakeParam struct {
		Title       string `db:"title,omitempty" json:"title"`
		Description string `db:"description,omitempty" json:"description"`
		Rating      *int   `db:"rating,omitempty" json:"rating"`
		Image       string `db:"image,omitempty" json:"image"`
	}

	// CakeDetail detail response
	CakeDetail struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Rating      *int   `json:"rating"`
		Image       string `json:"image"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
)
