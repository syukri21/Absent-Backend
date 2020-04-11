package model

// PStudentSchedule ...
type PStudentSchedule struct {
	StudentID  uint `json:"studentId"  gorm:"primary_key;auto_increment:false"`
	ScheduleID uint `json:"scheduleId"  `
	Semester   int  `json:"semester" gorm:"primary_key;auto_increment:false"`
	CourseID   uint `json:"courseId" gorm:"primary_key;auto_increment:false"`

	Student Student `gorm:"foreignkey:StudentID;association_foreignkey:UserID" json:"student"`
	Course  Course  `json:"course"`
}
