package model

// StudentSchedule ...
type StudentSchedule struct {
	StudentID  uint `json:"studentId"  gorm:"primary_key;auto_increment:false"`
	ScheduleID uint `json:"scheduleId"  `
	Semester   int  `json:"semester" gorm:"primary_key;auto_increment:false"`
	CourseID   uint `json:"courseId" gorm:"primary_key;auto_increment:false"`
}

// ShowStudentSchedule ...
type ShowStudentSchedule struct {
	StudentID  uint     `json:"studentId"  gorm:"primary_key;auto_increment:false"`
	ScheduleID uint     `json:"scheduleId"`
	Semester   int      `json:"semester" gorm:"primary_key;auto_increment:false"`
	CourseID   uint     `json:"courseId" gorm:"primary_key;auto_increment:false"`
	Course     *Course  `json:"course"`
	Student    *Student `gorm:"foreignkey:StudentID;association_foreignkey:UserID" json:"student"`
	Absent     *Absent  `gorm:"foreignkey:StudentID;association_foreignkey:StudentID"`
}

// TableName ...
func (*ShowStudentSchedule) TableName() string {
	return "student_schedules"
}
