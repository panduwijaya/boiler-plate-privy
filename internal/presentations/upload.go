// Package presentations
// Automatic generated
package presentations

type (
	// CakeQuery parameter
	UploadParam struct {
		ID          int    `db:"id,omitempty" json:"id" url:"id,omitempty"`
		Title       string `db:"title,omitempty" json:"title" url:"title,omitempty"`
		Rating      *int   `db:"rating,omitempty" json:"rating"`
		Description string `db:"description,omitempty" json:"description" url:"description,omitempty"`
		Image       string `db:"image,omitempty" json:"image" url:"image,omitempty"`
		Paging
		PeriodRange
	}
)
