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
	db.Model(&teacher.Teacher{}).AddForeignKey("user_id", "absents(teacher_id)", "CASCADE", "CASCADE")

	db.Model(&student.Student{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&student.Student{}).AddForeignKey("user_id", "absents(student_id)", "CASCADE", "CASCADE")

	fullname := "Syukri Husaibatul Khairi"
	nid := "1234567890"

	db.Debug().Create(&teacher.Teacher{
		Fullname: &fullname,
		Nid:      &nid,
		UserID:   1,
		User: user.User{
			Username: "ukiuki",
			RoleID:   1,
			Model: gorm.Model{
				ID: 1,
			},
		},
	})

	db.Debug().Create(&student.Student{
		Fullname: "Fuzi Widi",
		Nim:      "0001111",
		UserID:   2,
		User: user.User{
			Username: "fuziwidi",
			RoleID:   2,
			Model: gorm.Model{
				ID: 2,
			},
		},
	})

}
