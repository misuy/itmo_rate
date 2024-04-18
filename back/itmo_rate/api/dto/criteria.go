package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"
)

type CriteriaDTO struct {
	Name   string
	Rating float32
}

func CriteriaDTOFromCriteria(criteria entities.Criteria) CriteriaDTO {
	return CriteriaDTO{
		Name:   criteria.Name,
		Rating: criteria.Value,
	}
}

func CriteriaDTOListFromCriteriaList(list *entities.CriteriaList) []CriteriaDTO {
	return util.Map(list.List, CriteriaDTOFromCriteria)
}
