package roles

import "time"

type Role struct {
	ID          int       `json:"idrole" gorm:"primaryKey;autoIncrement"`
	Description string    `json:"descriptionRole"`
	Type        string    `json:"typeRole"    validate:"required,eq=ADMIN|eq=USER"`
	CreatedAt   time.Time `json:"Role_created_at"`
	
}
