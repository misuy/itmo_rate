package main

import (
	"itmo_rate/DB"
	"itmo_rate/DB/entities"
)

/*
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

type ObjectPreview struct {
	Id    int
	Name  string
	Tags  []string
	Score float32
}

type Review struct {
	Id       int
	Subject  string
	Lecturer string
	Teacher  string
	Text     string
	Author   string
	Criteria []Criteria
	Created  time.Time
}

type ReviewPreview struct {
	Id        int
	AvgRating float32
	Text      string
	Author    string
	Created   time.Time
}

var criteria Criteria = Criteria{"Unknown", 2}

var subjects []Subject = []Subject{
	{0, "ТПО", []string{"ВТ", "ФПИиКТ"}, []string{"Клименков Сергей Викторович"}, []string{"Клименков Сергей Викторович", "Соснов Николай Федорович"}, []Criteria{criteria, criteria, criteria, criteria, criteria}, 2.3},
	{1, "ОПД", []string{"ВТ", "ФПИиКТ"}, []string{"Клименков Сергей Викторович"}, []string{"Клименков Сергей Викторович"}, []Criteria{criteria, criteria, criteria, criteria, criteria}, 8.9},
}

var teachers []Teacher = []Teacher{
	{0, "Клименков Сергей Викторович", "path/to/image", []string{"ТПО", "ОПД"}, []Criteria{criteria, criteria, criteria, criteria, criteria}, 7.7},
	{1, "Соснов Николай Федорович", "path/to/image", []string{"ТПО"}, []Criteria{criteria, criteria, criteria, criteria, criteria}, 5.6},
}

var reviews []Review = []Review{
	{0, "ТПО", "Клименков Сергей Викторович", "Соснов Николай Федорович", "sample review text", "anon", []Criteria{criteria, criteria, criteria, criteria, criteria}, time.Now()},
	{1, "ТПО", "Клименков Сергей Викторович", "Клименков Сергей Викторович", "sample review text", "anon", []Criteria{criteria, criteria, criteria, criteria, criteria}, time.Now()},
	{2, "ОПД", "Клименков Сергей Викторович", "Клименков Сергей Викторович", "sample review text", "anon", []Criteria{criteria, criteria, criteria, criteria, criteria}, time.Now()},
	{3, "ОПД", "Клименков Сергей Викторович", "Клименков Сергей Викторович", "sample review text", "anon", []Criteria{criteria, criteria, criteria, criteria, criteria}, time.Now()},
}

func getSubjectById(id int) (subject Subject, ok bool) {
	if filtered := util.Filter(
		subjects,
		func(s Subject) bool {
			return s.Id == id
		},
	); len(filtered) == 0 {
		ok = false
	} else {
		subject = filtered[0]
		ok = true
	}
	return
}

func getTeacherById(id int) (teacher Teacher, ok bool) {
	if filtered := util.Filter(
		teachers,
		func(t Teacher) bool {
			return t.Id == id
		},
	); len(filtered) == 0 {
		ok = false
	} else {
		teacher = filtered[0]
		ok = true
	}
	return
}

func getReviewById(id int) (review Review, ok bool) {
	if filtered := util.Filter(
		reviews,
		func(r Review) bool {
			return r.Id == id
		},
	); len(filtered) == 0 {
		ok = false
	} else {
		review = filtered[0]
		ok = true
	}
	return
}

func (subject Subject) getReviews() []Review {
	return util.Filter(reviews,
		func(r Review) bool {
			return r.Subject == subject.Name
		},
	)
}

func (teacher Teacher) getReviews() []Review {
	return util.Filter(reviews,
		func(r Review) bool {
			return (r.Lecturer == teacher.Name) || (r.Teacher == teacher.Name)
		},
	)
}

func (subject Subject) toPreview() ObjectPreview {
	return ObjectPreview{subject.Id, subject.Name, subject.Faculties, subject.AvgRating}
}

func (teacher Teacher) toPreview() ObjectPreview {
	return ObjectPreview{teacher.Id, teacher.Name, teacher.Subjects, teacher.AvgRating}
}

func (review Review) toPreview() ReviewPreview {
	return ReviewPreview{
		review.Id,
		util.Mean(util.Map(
			review.Criteria,
			func(c Criteria) float32 {
				return c.Rating
			},
		)),
		review.Text,
		review.Author,
		review.Created,
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

	t := []ObjectPreview{}
	s := []ObjectPreview{}

	switch searchType {
	case "teachers":
		t = util.Map(
			util.Filter(
				teachers,
				func(teacher Teacher) bool {
					return strings.Contains(teacher.Name, text)
				},
			),
			func(teacher Teacher) ObjectPreview {
				return teacher.toPreview()
			},
		)
	case "subjects":
		s = util.Map(
			util.Filter(subjects,
				func(subject Subject) bool {
					return strings.Contains(subject.Name, text)
				},
			),
			func(subject Subject) ObjectPreview {
				return subject.toPreview()
			},
		)
	case "both":
		t = util.Map(
			util.Filter(
				teachers,
				func(teacher Teacher) bool {
					return strings.Contains(teacher.Name, text)
				},
			),
			func(teacher Teacher) ObjectPreview {
				return teacher.toPreview()
			},
		)
		s = util.Map(
			util.Filter(subjects,
				func(subject Subject) bool {
					return strings.Contains(subject.Name, text)
				},
			),
			func(subject Subject) ObjectPreview {
				return subject.toPreview()
			},
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

	ctx.JSON(http.StatusOK, util.Map(teachers[min(max(0, offset), len(teachers)):min(max(0, offset+amount), len(teachers))], func(teacher Teacher) ObjectPreview { return teacher.toPreview() }))
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

	ctx.JSON(http.StatusOK, util.Map(subjects[min(max(0, offset), len(subjects)):min(max(0, offset+amount), len(subjects))], func(subject Subject) ObjectPreview { return subject.toPreview() }))
}

func subjectEndpoint(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	subject, ok := getSubjectById(id)

	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, subject)
}

func teacherEndpoint(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	teacher, ok := getTeacherById(id)

	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

func subjectReviewsEndpoint(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

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

	subject, ok := getSubjectById(id)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	subjectReviews := subject.getReviews()

	ctx.JSON(http.StatusOK, util.Map(subjectReviews[min(max(0, offset), len(subjectReviews)):min(max(0, offset+amount), len(subjectReviews))], func(review Review) ReviewPreview { return review.toPreview() }))
}

func teacherReviewsEndpoint(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

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

	teacher, ok := getTeacherById(id)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	teacherReviews := teacher.getReviews()

	ctx.JSON(http.StatusOK, util.Map(teacherReviews[min(max(0, offset), len(teacherReviews)):min(max(0, offset+amount), len(teacherReviews))], func(review Review) ReviewPreview { return review.toPreview() }))
}

func reviewEndpoint(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	review, ok := getReviewById(id)

	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, review)

}
*/

func main() {
	/*
		router := gin.Default()

		api := router.Group("/api")
		api.GET("/user", userEndpoint)
		api.GET("/search", searchEndpoint)
		api.GET("/teachers", teachersEndpoint)
		api.GET("/subjects", subjectsEndpoint)
		api.GET("/subject/:id", subjectEndpoint)
		api.GET("/teacher/:id", teacherEndpoint)
		api.GET("/subject/:id/reviews", subjectReviewsEndpoint)
		api.GET("/teacher/:id/reviews", teacherReviewsEndpoint)
		api.GET("/review/:id", reviewEndpoint)

		router.Run(":8888")
	*/
	db, err := DB.Open()
	if err != nil {
		return
	}

	//DB.SaveTestData(db)

	var teachers []entities.Teacher
	var subjects []entities.Subject
	var reviews []entities.Review
	db.Model(&entities.Subject{}).Find(&subjects)
	db.Model(&entities.Review{}).Find(&reviews)
	println(len(reviews))
	//println(subjects[0].Teachers[0].Name)
	teachers, _ = reviews[0].GetTeachersByRole(db, "lecturer")
	println(len(teachers))
}
