package main

import (
	"backend-qrcode/model"
	"fmt"
	"strconv"
	"time"

	"log"

	fake "github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
)

// Migrate ...
func Migrate(db *gorm.DB) {

	fmt.Printf("Loading...")

	db.Debug().DropTableIfExists(
		&model.Grade{},
		&model.StudentSchedule{},
		&model.Absent{},
		&model.Schedule{},
		&model.Course{},
		&model.Teacher{},
		&model.Student{},
		&model.Admin{},
		&model.User{},
	)

	err := db.Debug().AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Teacher{},
		&model.Student{},
		&model.Course{},
		&model.Schedule{},
		&model.Absent{},
		&model.StudentSchedule{},
		&model.Grade{},
	).Error

	if err != nil {
		log.Fatal("Error Migration", err)
	}

	db.Debug().Model(&model.Teacher{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Debug().Model(&model.Student{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Debug().Model(&model.Admin{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Debug().Model(&model.Absent{}).AddForeignKey("student_id", "students(user_id)", "CASCADE", "CASCADE")
	db.Debug().Model(&model.Absent{}).AddForeignKey("teacher_id", "teachers(user_id)", "CASCADE", "CASCADE")
	db.Debug().Model(&model.Absent{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&model.Absent{}).AddForeignKey("schedule_id", "schedules(id)", "CASCADE", "CASCADE")

	db.Debug().Model(&model.Schedule{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&model.Schedule{}).AddForeignKey("teacher_id", "teachers(user_id)", "CASCADE", "CASCADE")

	db.Debug().Model(&model.StudentSchedule{}).AddForeignKey("student_id", "students(user_id)", "CASCADE", "CASCADE")
	db.Debug().Model(&model.StudentSchedule{}).AddForeignKey("schedule_id", "schedules(id)", "CASCADE", "CASCADE")

	db.Debug().Model(&model.Grade{}).AddForeignKey("course_id", "courses(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&model.Grade{}).AddForeignKey("schedule_id", "schedules(id)", "CASCADE", "CASCADE")
	db.Debug().Model(&model.Grade{}).AddForeignKey("student_id", "students(user_id)", "CASCADE", "CASCADE")

	fullname := "Syukri Husaibatul Khairi"
	nid := "1234567890"

	db.Debug().FirstOrCreate(&model.Admin{
		Fullname: fullname,
		NIA:      nid,
		UserID:   2,
		User: model.User{
			Username: "admin",
			RoleID:   3,
			Model: gorm.Model{
				ID: 2,
			},
			Hash: "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	db.Debug().FirstOrCreate(&model.Teacher{
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

	db.Debug().FirstOrCreate(&model.Course{
		Name:     "Kalkulus 1",
		Semester: 1,
		TotalSks: 3,
	})

	db.Debug().FirstOrCreate(&model.Schedule{
		ID:              1,
		CourseID:        1,
		TeacherID:       1,
		Day:             1,
		Week:            "BOTH",
		Time:            200,
		NumberOfMeeting: 1,
	})

	db.Debug().Create(&model.Student{
		Fullname: "Fuzi Widiastuti",
		Nim:      strconv.Itoa(int(time.Now().Unix()) - 1),
		UserID:   uint(3),
		User: model.User{
			Username: "fuziwidi123",
			RoleID:   2,
			Model:    gorm.Model{ID: uint(3)},
			Hash:     "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
		},
	})

	// Makse Student
	for i := 0; i < 29; i++ {
		db.Debug().Create(&model.Student{
			Fullname: fake.FirstName() + " " + fake.LastName(),
			Nim:      strconv.Itoa(int(time.Now().Unix()) + i),
			UserID:   uint(i + 4),
			User: model.User{
				Username: fake.UserName(),
				RoleID:   2,
				Model:    gorm.Model{ID: uint(i + 4)},
				Hash:     "$2a$04$A75O8a8W2Ze1LwX4oY0UB.B6xwHsQlPRc66vbBnPMcQs28S7hsWWG",
			},
		})
	}

	// Register Schedule
	for i := 0; i < 30; i++ {
		db.Debug().Create(&model.StudentSchedule{
			ScheduleID: 1,
			CourseID:   1,
			Semester:   1,
			StudentID:  uint(i + 3),
		})
	}

	fmt.Printf("Done...")

}
