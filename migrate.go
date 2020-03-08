package main

import (
	absent "backend-qrcode/absent/handler"
	course "backend-qrcode/course/handler"
	schedule "backend-qrcode/schedule/handler"
	student "backend-qrcode/student/handler"
	teacher "backend-qrcode/teacher/handler"
	user "backend-qrcode/user/handler"

	"log"

	"github.com/jinzhu/gorm"
)

// Migrate ...
func Migrate(db *gorm.DB) {

	db.DropTableIfExists(
		&schedule.Schedule{},
		&course.Course{},
		&absent.Absent{},
		&teacher.Teacher{},
		&student.Student{},
		&user.User{},
	)

	err := db.AutoMigrate(
		&user.User{},
		&teacher.Teacher{},
		&student.Student{},
		&course.Course{},
		&absent.Absent{},
		&schedule.Schedule{},
	).Error

	if err != nil {
		log.Fatal("Error Migration", err)
	}

	db.Model(&teacher.Teacher{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&student.Student{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Model(&absent.Absent{}).AddForeignKey("student_id", "students(user_id)", "CASCADE", "CASCADE")
	db.Model(&absent.Absent{}).AddForeignKey("teacher_id", "teachers(user_id)", "CASCADE", "CASCADE")
	db.Model(&absent.Absent{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")
	db.Model(&absent.Absent{}).AddForeignKey("schedule_id", "schedules(id)", "CASCADE", "CASCADE")

	db.Model(&schedule.Schedule{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")
	db.Model(&schedule.Schedule{}).AddForeignKey("teacher_id", "teachers(user_id)", "CASCADE", "CASCADE")

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
		Fullname: "Tuuuut",
		Nim:      "0001111",
		UserID:   2,
		User: user.User{
			Username: "tutituti",
			RoleID:   2,
			Model: gorm.Model{
				ID: 2,
			},
			Hash: "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	db.Debug().FirstOrCreate(&course.Course{
		Name:     "Kalkulus 1",
		Semester: 1,
		TotalSks: 3,
	})

	db.Debug().FirstOrCreate(&schedule.Schedule{
		CourseID:  1,
		TeacherID: 1,
		Day:       1,
		Week:      "BOTH",
		Time:      200,
	})

}
