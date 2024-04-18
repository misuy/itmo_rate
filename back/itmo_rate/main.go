package main

import (
	"itmo_rate/DB"
	"itmo_rate/api"
)

func main() {
	db, err := DB.Open()
	if err != nil {
		return
	}
	//DB.SaveTestData(db)

	api.Run(db, 8888)
}
