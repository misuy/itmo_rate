package entities

import "gorm.io/gorm"

type Criteria struct {
	gorm.Model
	Name           string
	Value          float32
	CriteriaListID uint
}

func NewCriteria(name string, value float32) Criteria {
	return Criteria{
		Name:  name,
		Value: value,
	}
}

type CriteriaList struct {
	gorm.Model
	List      []Criteria
	ReviewID  *uint
	SubjectID *uint
	TeacherID *uint
}

func NewCriteriaList(list []Criteria) CriteriaList {
	return CriteriaList{
		List: list,
	}
}

type DefaultCriteriaList struct {
	Criteria1Name  string
	Criteria1Value float32

	Criteria2Name  string
	Criteria2Value float32

	Criteria3Name  string
	Criteria3Value float32

	Criteria4Name  string
	Criteria4Value float32

	Criteria5Name  string
	Criteria5Value float32
}

func NewDefaultCriteriaList(criteria1Value float32, criteria2Value float32, criteria3Value float32, criteria4Value float32, criteria5Value float32) DefaultCriteriaList {
	return DefaultCriteriaList{
		Criteria1Name:  "Критерий 1",
		Criteria1Value: criteria1Value,

		Criteria2Name:  "Критерий 2",
		Criteria2Value: criteria2Value,

		Criteria3Name:  "Критерий 3",
		Criteria3Value: criteria3Value,

		Criteria4Name:  "Критерий 4",
		Criteria4Value: criteria4Value,

		Criteria5Name:  "Критерий 5",
		Criteria5Value: criteria5Value,
	}
}

func (defaultCriteriaList DefaultCriteriaList) ToCriteriaList() CriteriaList {
	return NewCriteriaList(
		[]Criteria{
			NewCriteria(defaultCriteriaList.Criteria1Name, defaultCriteriaList.Criteria1Value),
			NewCriteria(defaultCriteriaList.Criteria2Name, defaultCriteriaList.Criteria2Value),
			NewCriteria(defaultCriteriaList.Criteria3Name, defaultCriteriaList.Criteria3Value),
			NewCriteria(defaultCriteriaList.Criteria4Name, defaultCriteriaList.Criteria4Value),
			NewCriteria(defaultCriteriaList.Criteria5Name, defaultCriteriaList.Criteria5Value),
		},
	)
}
