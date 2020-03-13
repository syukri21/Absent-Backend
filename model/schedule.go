package model

import "time"

// Schedule ...
type Schedule struct {
	ID        uint    `json:"id" gorm:"not null;unique_index;AUTO_INCREMENT"`
	CourseID  uint    `json:"courseId" gorm:"primary_key;auto_increment:false"`
	TeacherID uint    `json:"teacherId" gorm:"primary_key;auto_increment:false"`
	Day       int     `json:"day"`
	Week      string  `json:"week"`
	Time      int     `json:"time"`
	Teacher   Teacher `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course    Course
	Absents   []Absent `gorm:"foreignkey:ScheduleID"`
}

// Model ...
type ScheduleModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

/* -------------------------------------------------------------------------- */
/*                                 NOTE INDEX                                 */
/* -------------------------------------------------------------------------- */

type ScheduleStudent struct {
	UserID   uint   `json:"userId"`
	Nim      string `json:"nim"`
	Fullname string `json:"fullname"`
}

type ScheduleAbsent struct {
	ScheduleID       uint            `json:"scheduleId" gorm:"primary_key;auto_increment:false"`
	AbsentHash       string          `json:"-" gorm:"unique_index"`
	StudentID        uint            `json:"studentId" gorm:"primary_key;auto_increment:false"`
	NumberOfMeetings int             `json:"numberOfMeetings" `
	Semester         int             `json:"semester"`
	AbsentTime       *time.Time      `json:"absentTime" `
	Student          ScheduleStudent `gorm:"foreignkey:StudentID;association_foreignkey:UserID"`
	ScheduleModel
}

type ScheduleIndex struct {
	ID        uint   `json:"id"`
	CourseID  uint   `json:"courseId"`
	TeacherID uint   `json:"teacherId"`
	Day       int    `json:"day"`
	Week      string `json:"week"`
	Time      int    `json:"time"`
	Course    Course
}

func (ScheduleIndex) TableName() string {
	return "schedules"
}

/* -------------------------------------------------------------------------- */
/*                              NOTE SCHEDULESHOW                             */
/* -------------------------------------------------------------------------- */

type ScheduleShow struct {
	ScheduleIndex
	Absents []Absent `gorm:"foreignkey:ID;association_foreignkey:ScheduleID"`
}

func (ScheduleShow) TableName() string {
	return "schedules"
}

/* -------------------------------------------------------------------------- */
/*                                 NOTE CREATE                                */
/* -------------------------------------------------------------------------- */

// CreateParams ...
type ScheduleCreateParams struct {
	CourseID uint   `json:"courseId"`
	Day      int    `json:"day"`
	Week     string `json:"week"`
	Time     int    `json:"time"`
}

// CreateSchedule ...
type ScheduleCreate struct {
	ID        uint   `json:"id"`
	CourseID  uint   `json:"courseId"`
	TeacherID uint   `json:"teacherId"`
	Day       int    `json:"day"`
	Week      string `json:"week"`
	Time      int    `json:"time"`
}

// TableName ...
func (ScheduleCreate) TableName() string {
	return "schedules"

}
