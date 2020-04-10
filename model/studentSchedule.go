package model

// PStudentSchedule ...
type PStudentSchedule struct {
	StudentID  int `json:"studentId"  gorm:"primary_key;auto_increment:false"`
	ScheduleID int `json:"scheduleId"`
	Semester   int `json:"semester" gorm:"primary_key;auto_increment:false"`
	CourseID   int `json:"courseId" gorm:"primary_key;auto_increment:false"`
}
