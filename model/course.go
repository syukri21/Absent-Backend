package model

// Course ...
type Course struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"name" gorm:"unique_index" `
	TotalSks int    `json:"totalSks"`
	Semester int    `json:"semester"`
}

// CourseDeleted ...
type CourseDeleted struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

// CourseEditParams ...
type CourseEditParams struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	TotalSks int    `json:"totalSks"`
	Semester int    `json:"semester"`
}
