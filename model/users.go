package model

type User struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Point     int64  `json:"point" db:"point"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
	DeletedAt *int64 `json:"deletedAt,omitempty" db:"deleted_at"`
}
