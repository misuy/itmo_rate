package dto

import (
	"itmo_rate/DB/entities"
	"itmo_rate/util"
)

type CriteriaDTO struct {
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
}

func CriteriaDTOFromCriteria(criteria entities.Criteria) CriteriaDTO {
	return CriteriaDTO{
		Name:   criteria.Name,
		Rating: criteria.Value,
	}
}

func CriteriaDTOToCriteria(criteriaDTO CriteriaDTO) entities.Criteria {
	return entities.NewCriteria(criteriaDTO.Name, criteriaDTO.Rating)
}

func CriteriaDTOListFromCriteriaList(list *entities.CriteriaList) []CriteriaDTO {
	return util.Map(list.List, CriteriaDTOFromCriteria)
}

func CriteriaDTOListToCriteriaList(list []CriteriaDTO) entities.CriteriaList {
	criteriaList := util.Map(list, CriteriaDTOToCriteria)
	return entities.NewCriteriaList(criteriaList)
}
