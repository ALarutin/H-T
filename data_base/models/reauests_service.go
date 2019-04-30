package models

func (db *dbManager) ClearDatabase() (err error) {
	_, err = db.dataBase.Exec(`SELECT clear_database()`)
	return
}
