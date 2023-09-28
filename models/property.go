package models

import "time"

type Property struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"type:varchar(100)" json:"property_type"`
	Address   string    `gorm:"type:varchar(200)" json:"address"`
	Price     int       `gorm:"type:integer" json:"price_rental"`
	Status    string    `gorm:"type:varchar(50)" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PropertyTenantResponse struct {
	ID      uint   `json:"id"`
	Type    string `json:"property_type"`
	Address string `json:"address"`
	Price   int    `json:"price_rental"`
	Status  string `json:"status"`
}
