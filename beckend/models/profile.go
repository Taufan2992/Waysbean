package models

import "time"

type Profile struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Address   string               `json:"address" gorm:"type: text"`
	Postcode  string               `json:"postcode" gorm:"type: text"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type ProfileResponse struct {
	ID       int                  `json:"id"`
	Address  string               `json:"address"`
	Postcode string               `json:"postcode"`
	UserID   int                  `json:"user_id"`
	User     UsersProfileResponse `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
