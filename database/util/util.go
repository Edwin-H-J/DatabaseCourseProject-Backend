package util

import "backend/database/internel"
import "database/sql"

func StartTx() (*sql.Tx,error){
	return database.Db.Begin()
}