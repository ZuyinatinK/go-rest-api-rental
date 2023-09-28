package models

import "time"

type Tenant struct {
	ID         uint      `grom:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Contact    int       `json:"contact"`
	PropertyID uint      `json:"property_id"`
	Property   Property  `gorm:"foreignKey:PropertyID" json:"property"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TenantResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Contact    int       `json:"contact"`
	PropertyID uint      `json:"-"`
	Property   Property  `gorm:"foreignKey:PropertyID" json:"property"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
