package dto

import "itmo_rate/DB/entities"

type UserDTO struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Id   uint   `json:"id"`
}

func UserDTOFromUser(user *entities.User) UserDTO {
	return UserDTO{
		Name: user.Name,
		Role: user.Role,
		Id:   user.ID,
	}
}
