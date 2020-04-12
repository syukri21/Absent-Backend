package model

// Grade ...
type Grade struct {
	StudentID   uint    `json:"studentId"  gorm:"primary_key;auto_increment:false"`
	ScheduleID  uint    `json:"scheduleId" `
	Semester    int     `json:"semester" gorm:"primary_key;auto_increment:false"`
	CourseID    uint    `json:"courseId" gorm:"primary_key;auto_increment:false"`
	Attendance  float32 `json:"attendance" gorm:"type:DECIMAL(3,2)"`
	Assignment  float32 `json:"assignment" gorm:"type:DECIMAL(3,2)"`
	Uts         float32 `json:"uts" gorm:"type:DECIMAL(3,2)"`
	Uas         float32 `json:"uas" gorm:"type:DECIMAL(3,2)"`
	WeightValue float32 `json:"weightValue" gorm:"type:DECIMAL(3,2)"`
	LetterValue string  `json:"letterValue" gorm:"type:varchar(4)"`
}

// ShowGradeEntity ...
type ShowGradeEntity struct {
	StudentID   uint    `json:"-"  gorm:"primary_key;auto_increment:false"`
	ScheduleID  uint    `json:"-" `
	Semester    int     `json:"-" gorm:"primary_key;auto_increment:false"`
	CourseID    uint    `json:"-" gorm:"primary_key;auto_increment:false"`
	Attendance  float32 `json:"attendance" gorm:"type:DECIMAL(3,2)"`
	Assignment  float32 `json:"assignment" gorm:"type:DECIMAL(3,2)"`
	Uts         float32 `json:"uts" gorm:"type:DECIMAL(3,2)"`
	Uas         float32 `json:"uas" gorm:"type:DECIMAL(3,2)"`
	WeightValue float32 `json:"weightValue" gorm:"type:DECIMAL(3,2)"`
	LetterValue string  `json:"letterValue" gorm:"type:varchar(4)"`
}

// TableName ...
func (*ShowGradeEntity) TableName() string {
	return "grades"
}

// ShowGradeByScheduleID ...
type ShowGradeByScheduleID struct {
	ShowGradeEntity
	Grade   *ShowGradeEntity `json:"grade" gorm:"foreignkey:StudentID,ScheduleID;association_foreignkey:StudentID,ScheduleID" `
	Student *Student         `gorm:"foreignkey:StudentID;association_foreignkey:UserID" json:"student"`
}

// TableName ...
func (*ShowGradeByScheduleID) TableName() string {
	return "student_schedules"
}
