package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"

	"strings"

	"itmo_rate/util"
)

type User struct {
	Name string
	Role string
}

type Criteria struct {
	Name   string
	Rating float32
}

type Subject struct {
	Id        int
	Name      string
	Faculties []string
	Lecturers []string
	Teachers  []string
	Criteria  []Criteria
	AvgRating float32
}

type Teacher struct {
	Id        int
	Name      string
	Avatar    string
	Subjects  []string
	Criteria  []Criteria
	AvgRating float32
}

type ShortObject struct {
	Id    int
	Name  string
	Tags  []string
	Score float32
}

var criteria Criteria = Criteria{"Unknown", 2}

var subjects []Subject = []Subject{
	{0, "ТПО", []string{"ВТ", "ФПИиКТ"}, []string{"Клименков Сергей Викторович"}, []string{"Клименков Сергей Викторович", "Соснов Николай Федорович"}, []Criteria{criteria, criteria, criteria, criteria}, 2.3},
	{1, "ОПД", []string{"ВТ", "ФПИиКТ"}, []string{"Клименков Сергей Викторович"}, []string{"Клименков Сергей Викторович"}, []Criteria{criteria, criteria, criteria, criteria}, 8.9},
}

var teachers []Teacher = []Teacher{
	{0, "Клименков Сергей Викторович", "path/to/image", []string{"ТПО", "ОПД"}, []Criteria{criteria, criteria, criteria, criteria}, 7.7},
	{1, "Соснов Николай Федорович", "path/to/image", []string{"ТПО"}, []Criteria{criteria, criteria, criteria, criteria}, 5.6},
}

func subjectToShortObject(subject Subject) ShortObject {
	return ShortObject{subject.Id, subject.Name, subject.Faculties, subject.AvgRating}
}

func teacherToShortObject(teacher Teacher) ShortObject {
	return ShortObject{
		teacher.Id,
		teacher.Name,
		teacher.Subjects,
		teacher.AvgRating,
	}
}

func userEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, User{"Misha", "admin"})
}

func searchEndpoint(ctx *gin.Context) {
	searchType, ok := ctx.GetQuery("type")
	if !ok {
		searchType = "both"
	}

	text, ok := ctx.GetQuery("text")
	if !ok {
		text = ""
	}

	t := []ShortObject{}
	s := []ShortObject{}

	switch searchType {
	case "teachers":
		t = util.Map(
			util.Filter(teachers,
				func(teacher Teacher) bool {
					return strings.Contains(teacher.Name, text)
				},
			),
			teacherToShortObject,
		)
	case "subjects":
		s = util.Map(
			util.Filter(subjects,
				func(subject Subject) bool {
					return strings.Contains(subject.Name, text)
				},
			),
			subjectToShortObject,
		)
	case "both":
		t = util.Map(
			util.Filter(teachers,
				func(teacher Teacher) bool {
					return strings.Contains(teacher.Name, text)
				},
			),
			teacherToShortObject,
		)
		s = util.Map(
			util.Filter(subjects,
				func(subject Subject) bool {
					return strings.Contains(subject.Name, text)
				},
			),
			subjectToShortObject,
		)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"teachers": t,
		"subjects": s,
	})
}

func teachersEndpoint(ctx *gin.Context) {
	offset := 0
	if offsett, ok := ctx.GetQuery("offset"); ok {
		if offsett, err := strconv.Atoi(offsett); err == nil {
			offset = offsett
		}
	}

	amount := 0
	if amountt, ok := ctx.GetQuery("amount"); ok {
		if amountt, err := strconv.Atoi(amountt); err == nil {
			amount = amountt
		}
	}

	ctx.JSON(http.StatusOK, util.Map(teachers[min(max(0, offset), len(teachers)):min(max(0, offset+amount), len(teachers))], teacherToShortObject))
}

func subjectsEndpoint(ctx *gin.Context) {
	offset := 0
	if offsett, ok := ctx.GetQuery("offset"); ok {
		if offsett, err := strconv.Atoi(offsett); err == nil {
			offset = offsett
		}
	}

	amount := 0
	if amountt, ok := ctx.GetQuery("amount"); ok {
		if amountt, err := strconv.Atoi(amountt); err == nil {
			amount = amountt
		}
	}

	ctx.JSON(http.StatusOK, util.Map(subjects[min(max(0, offset), len(subjects)):min(max(0, offset+amount), len(subjects))], subjectToShortObject))
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/user", userEndpoint)
	api.GET("/search", searchEndpoint)
	api.GET("/teachers", teachersEndpoint)
	api.GET("/subjects", subjectsEndpoint)

	router.Run(":8888")
}
