package models

type GameRecord struct {
	ID                uint   `json:"id" gorm:"primary_key"`
	Name              string `json:"name"`
	Platform          string `json:"platform"` // i.e "PC" | "PSP" etc..
	Medal             string `json:"medal"`    // "Gold" | "Silver"
	CompleteTimeHours int    `json:"complete_time_hours"`
	Genre             string `json:"genre"`
}
