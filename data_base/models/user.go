package models

import (
	"data_base/presentation/logger"
)

func (db *dbManager) GetUser (nickname string)(user User, err error){

	row := db.dataBase.QueryRow(`SELECT nickname, email, fullname, about FROM public."person" WHERE nickname = $1`, nickname)
	err = row.Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	return
}

