package authdto

type RegisterResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
}

type LoginResponse struct {
	ID     int    `json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
	Image  string `gorm:"type: varchar(50)"  json:"image"`
	Token  string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
}
