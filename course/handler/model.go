package handler

// Course ...
type Course struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"name"`
	TotalSks int    `json:"totalSks"`
	Semester int    `json:"semeseter"`
}
