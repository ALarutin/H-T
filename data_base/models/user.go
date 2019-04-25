package models

func (db *dbManager) GetUser(nickname string) (user User, err error) {

	row := db.dataBase.QueryRow(
		`SELECT nickname, email, fullname, about FROM public."person" 
			WHERE nickname = $1`,
			nickname)
	err = row.Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
	return
}

func (db *dbManager) CreateUser(user User) (err error){

	_, err = db.dataBase.Exec(
		`INSERT INTO public."person" (email, about, fullname, nickname)
			  VALUES ($1, $2, $3, $4)`,
			  user.Email, user.About, user.Fullname, user.Nickname)
	return
}

func (db *dbManager) SelectUsers(nickname string, email string) (users []User, err error){

	rows, err := db.dataBase.Query(
		`SELECT nickname, email, fullname, about FROM public."person" 
			WHERE nickname = $1 OR email = $2`,
			nickname, email)
	if err !=  nil{
		return
	}

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

