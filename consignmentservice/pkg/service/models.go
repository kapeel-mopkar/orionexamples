package service

import "time"

type Consignment struct {
	ID          int64  `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Description string `json:"description"`
	Weight      int64  `gorm:"column:weight" json:"weight"`
	Containers  []Container
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type Container struct {
	ID            int64     `gorm:"column:id; PRIMARY_KEY" json:"id"`
	CustomerID    string    `gorm:"column:customer_id" json:"customer_id"`
	UserID        string    `gorm:"column:user_id" json:"user_id"`
	Origin        string    `gorm:"column:origin" json:"origin"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ConsignmentID int64     `gorm:"column:consignment_id"`
}
