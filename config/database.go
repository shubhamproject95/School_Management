package config

import (
	"School_gql/pkg/model"
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=school port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		_ = errors.New("can't connect to database")
		return nil
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(
		&model.Class{},
		&model.Complain{},
		&model.Exam{},
		&model.Fees{},
		&model.StudentAttendence{},
		&model.SessionYear{},
		&model.Student{},
		&model.Subject{},
		&model.TeacherAttendence{},
		&model.Teacher{},
		&model.Scholarship{},
		&model.Guardian{},
		&model.ClassTeacher{},
		&model.StudentClass{},
		&model.StudentExam{},
		&model.User{},
		&model.Role{},
		&model.ResetPassword{},
		&model.User{},
		&model.Login{},
		&model.Staff{},
		&model.Salary{},
		&model.BankDetail{},
		&model.StudentHomework{},
		&model.Guardian{},
		&model.Homework{},
		&model.ClassSubject{},
		&model.Tour{},
		&model.Result{},
	)
	if err != nil {
		log.Println(err.Error())
	}
}

//func SeedData(db *gorm.DB) {
//seeder := seeders.Init(db)

// seeder.SeedStaff()*
// seeder.SeedClass()*
// seeder.SeedBankDetails()*
// seeder.SeedStudent()*
// seeder.SeedStudentAttendene()*
// seeder.SeedTeacher()*
// seeder.SeedTeacherAttendence()*
// seeder.SeedSubject()*
// seeder.SeedRole()*
// seeder.SeedExam()*
// seeder.SeedComplain()*
// seeder.SeedFees()
// seeder.SeedSessionYear()*
// seeder.SeedClassSubject()*
//seeder.SeedExam()*
// seeder.SeedSalary()*
// seeder.SeedClassTeacher()*
// seeder.SeedGuardian()
// seeder.SeedHomework()
//seeder.SeedUser()*
// seeder.SeedUpdatePassword()
// seeder.SeedResetPassword()
// seeder.SeedResult()
// seeder.SeedStudentHomework()
//seeder.SeedStudentClass()*
// seeder.SeedStudentExam()
//seeder.SeedGuardian()*

//}
