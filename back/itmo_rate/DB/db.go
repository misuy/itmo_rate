package DB

import (
	"itmo_rate/DB/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (db *gorm.DB, err error) {
	dsn := "host=localhost user=itmo_rate dbname=itmo_rate_db port=5432"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(
		&entities.Faculty{},
		&entities.Subject{},
		&entities.Teacher{},
		&entities.User{},
		&entities.Review{},
		&entities.TeacherReviewRel{},
		&entities.SubjectTeacherRel{},
		&entities.CriteriaList{},
		&entities.Criteria{},
	)

	db.SetupJoinTable(&entities.Teacher{}, "Reviews", &entities.TeacherReviewRel{})
	db.SetupJoinTable(&entities.Subject{}, "Teachers", &entities.SubjectTeacherRel{})

	return
}

func SaveTestData(db *gorm.DB) {
	faculties := []entities.Faculty{
		entities.NewFaculty("ФПИиКТ"),
		entities.NewFaculty("ВТ"),
	}

	subjects := []entities.Subject{
		entities.NewSubject("ТПО"),
		entities.NewSubject("ОПД"),
	}

	teachers := []entities.Teacher{
		entities.NewTeacher("Клименков Сергей Викторович"),
		entities.NewTeacher("Соснов Георгий Александрович"),
	}

	users := []entities.User{
		entities.NewUser("Ненов Владислав"),
		entities.NewUser("Передрий Михаил"),
	}

	reviews := []entities.Review{
		entities.NewReview(entities.NewDefaultCriteriaList(5, 5, 5, 5, 5).ToCriteriaList(), "sample review text", true),
		entities.NewReview(entities.NewDefaultCriteriaList(7.8, 7.4, 2, 10, 4).ToCriteriaList(), "sample review text", true),
	}

	db.Create(&reviews)
	db.Create(&users)
	db.Create(&teachers)
	db.Create(&subjects)
	db.Create(&faculties)

	faculties[0].AddSubjects(db, []entities.Subject{subjects[0]})
	faculties[1].AddSubjects(db, subjects)

	subjects[0].AddTeacher(db, &teachers[0], "teacher")
	subjects[0].AddTeacher(db, &teachers[1], "lecturer")
	subjects[0].AddReviews(db, []entities.Review{reviews[1]})
	subjects[1].AddTeacher(db, &teachers[0], "lecturer")
	subjects[1].AddTeacher(db, &teachers[1], "teacher")
	subjects[1].AddReviews(db, []entities.Review{reviews[0]})

	teachers[0].AddReview(db, &reviews[0], "lecturer")
	teachers[0].AddReview(db, &reviews[1], "teacher")
	teachers[1].AddReview(db, &reviews[0], "teacher")
	teachers[1].AddReview(db, &reviews[1], "lecturer")

	users[0].AddReviews(db, []entities.Review{reviews[0]})
	users[1].AddReviews(db, []entities.Review{reviews[1]})

	db.Save(&reviews)
	db.Save(&users)
	db.Save(&teachers)
	db.Save(&subjects)
	db.Save(&faculties)
}
