package main

import (
	"backend-qrcode/model"

	"log"

	"github.com/jinzhu/gorm"
)

// Migrate ...
func Migrate(db *gorm.DB) {

	db.DropTableIfExists(
		&model.Schedule{},
		&model.Course{},
		&model.Absent{},
		&model.Teacher{},
		&model.Student{},
		&model.Admin{},
		&model.User{},
	)

	err := db.AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Teacher{},
		&model.Student{},
		&model.Course{},
		&model.Absent{},
		&model.Schedule{},
	).Error

	if err != nil {
		log.Fatal("Error Migration", err)
	}

	db.Model(&model.Teacher{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&model.Student{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&model.Admin{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&model.Absent{}).AddForeignKey("student_id", "students(user_id)", "CASCADE", "CASCADE")
	db.Model(&model.Absent{}).AddForeignKey("teacher_id", "teachers(user_id)", "CASCADE", "CASCADE")
	db.Model(&model.Absent{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")
	db.Model(&model.Absent{}).AddForeignKey("schedule_id", "schedules(id)", "CASCADE", "CASCADE")

	db.Model(&model.Schedule{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")
	db.Model(&model.Schedule{}).AddForeignKey("teacher_id", "teachers(user_id)", "CASCADE", "CASCADE")

	fullname := "Syukri Husaibatul Khairi"
	nid := "1234567890"

	db.FirstOrCreate(&model.Admin{
		Fullname: fullname,
		NIA:      nid,
		UserID:   3,
		User: model.User{
			Username: "admin",
			RoleID:   3,
			Model: gorm.Model{
				ID: 3,
			},
			Hash: "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	db.FirstOrCreate(&model.Teacher{
		Fullname: &fullname,
		Nid:      &nid,
		UserID:   1,
		User: model.User{
			Username: "ukiuki",
			RoleID:   1,
			Model: gorm.Model{
				ID: 1,
			},
			Hash: "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	db.FirstOrCreate(&model.Student{
		Fullname: "student01",
		Nim:      "0001111111",
		UserID:   2,
		User: model.User{
			Username: "tester",
			RoleID:   2,
			Model: gorm.Model{
				ID: 2,
			},
			Hash: "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	db.FirstOrCreate(&model.Course{
		Name:     "Kalkulus 1",
		Semester: 1,
		TotalSks: 3,
	})

	db.FirstOrCreate(&model.Schedule{
		ID:              1,
		CourseID:        1,
		TeacherID:       1,
		Day:             1,
		Week:            "BOTH",
		Time:            200,
		NumberOfMeeting: 1,
	})

}
