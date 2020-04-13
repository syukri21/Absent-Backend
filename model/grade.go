package model

// Grade ...
type Grade struct {
	ScheduleID  uint    `json:"scheduleId" `
	StudentID   uint    `json:"studentId" gorm:"primary_key;auto_increment:false"`
	Semester    int     `json:"semester" gorm:"primary_key;auto_increment:false"`
	CourseID    uint    `json:"courseId" gorm:"primary_key;auto_increment:false"`
	Attendance  float32 `json:"attendance" gorm:"type:DECIMAL(5,2)"`
	Assignment  float32 `json:"assignment" gorm:"type:DECIMAL(5,2)"`
	Uts         float32 `json:"uts" gorm:"type:DECIMAL(5,2)" valid:"required"`
	Uas         float32 `json:"uas" gorm:"type:DECIMAL(5,2)"`
	WeightValue float32 `json:"weightValue" gorm:"type:DECIMAL(5,2)"`
	LetterValue string  `json:"letterValue" gorm:"type:varchar(4)" valid:"required"`
}

// ShowGradeEntity ...
type ShowGradeEntity struct {
	ScheduleID  uint    `json:"-" `
	StudentID   uint    `json:"-" gorm:"primary_key;auto_increment:false"`
	Attendance  float32 `json:"attendance" gorm:"type:DECIMAL(5,2)"`
	Assignment  float32 `json:"assignment" gorm:"type:DECIMAL(5,2)"`
	Uts         float32 `json:"uts" gorm:"type:DECIMAL(5,2)"`
	Uas         float32 `json:"uas" gorm:"type:DECIMAL(5,2)"`
	WeightValue float32 `json:"weightValue" gorm:"type:DECIMAL(5,2)"`
	LetterValue string  `json:"letterValue" gorm:"type:varchar(4)"`
}

// TableName ...
func (*ShowGradeEntity) TableName() string {
	return "grades"
}

// ShowGradeByScheduleID ...
type ShowGradeByScheduleID struct {
	StudentSchedule
	Grade   *ShowGradeEntity `json:"grade" gorm:"foreignkey:StudentID,ScheduleID;association_foreignkey:StudentID,ScheduleID" `
	Student *Student         `json:"student" gorm:"foreignkey:StudentID;association_foreignkey:UserID" `
}

// TableName ...
func (*ShowGradeByScheduleID) TableName() string {
	return "student_schedules"
}
