package main

import (
	absent "backend-qrcode/absent/handler"
	course "backend-qrcode/course/handler"
	student "backend-qrcode/student/handler"
	teacher "backend-qrcode/teacher/handler"
	user "backend-qrcode/user/handler"

	"log"

	"github.com/jinzhu/gorm"
)

// Migrate ...
func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(&user.User{}, &teacher.Teacher{}, &student.Student{}, &course.Course{}, &absent.Absent{}).Error

	if err != nil {
		log.Fatal("Error Migration", err)
	}

	db.Model(&teacher.Teacher{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&student.Student{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&absent.Absent{}).AddForeignKey("student_id", "students(user_id)", "CASCADE", "CASCADE")
	db.Model(&absent.Absent{}).AddForeignKey("teacher_id", "teachers(user_id)", "CASCADE", "CASCADE")
	db.Model(&absent.Absent{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")

	fullname := "Syukri Husaibatul Khairi"
	nid := "1234567890"

	db.Debug().FirstOrCreate(&teacher.Teacher{
		Fullname: &fullname,
		Nid:      &nid,
		UserID:   1,
		User: user.User{
			Username: "ukiuki",
			RoleID:   1,
			Model: gorm.Model{
				ID: 1,
			},
			Hash: "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	db.Debug().FirstOrCreate(&student.Student{
		Fullname: "Fuzi Widi",
		Nim:      "0001111",
		UserID:   2,
		User: user.User{
			Username: "fuziwidi",
			RoleID:   2,
			Model: gorm.Model{
				ID: 2,
			},
			Hash: "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	db.Debug().FirstOrCreate(&course.Course{
		Name:     "Kalkulus",
		Semester: 1,
		TotalSks: 3,
	})

}
