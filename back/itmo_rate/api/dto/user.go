package dto

import "itmo_rate/DB/entities"

type UserDTO struct {
	Name string
	Role string
}

func UserDTOFromUser(user *entities.User) UserDTO {
	return UserDTO{
		Name: user.Name,
		Role: user.Role,
	}
}
