package user

import (
	"backend/database/internel"
	"database/sql"
	"errors"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Auth      string
	InfoId    int64
	Available int8
}

func GetUserById(uid int64, tx *sql.Tx) (User, error) {
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow("select * from user where id = ?", uid)
	} else {
		rows = database.Db.QueryRow("select * from user where id = ?", uid)
	}
	var user User
	err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Auth, &user.InfoId, &user.Available)
	return user, err
}
func GetAllUser(tx *sql.Tx) ([]User, error) {
	var err error
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.Query("select id,username,auth,info_id,available from user")
	} else {
		rows, err = database.Db.Query("select id,username,auth,info_id,available from user")
	}
	users := make([]User, 0)
	if err != nil {
		return users, err
	}
	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Auth, &user.InfoId, &user.Available)
		users = append(users, user)
		if err != nil {
			return users, err
		}
	}
	return users, nil
}
func GetUserByUsername(username string, tx *sql.Tx) (User, error) {
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow(`select * from user where username = ?;`, username)
	} else {
		rows = database.Db.QueryRow(`select * from user where username = ?;`, username)
	}
	var user User
	err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Auth, &user.InfoId, &user.Available)
	return user, err
}

func checkUsernameExist(username string) bool {
	rows := database.Db.QueryRow(`select count(*) from user where username = ?`, username)
	var num int
	rows.Scan(&num)
	if num == 0 {
		return false
	} else {
		return true
	}
}

func CreateUser(user User, tx *sql.Tx) (User, error) {
	if checkUsernameExist(user.Username) {
		return user, errors.New("exist same username")
	}
	var ret sql.Result
	var err error
	if tx != nil {
		ret, err = tx.Exec(`insert into user (username,password,auth,info_id,available) values(?,?,?,?,?)`, user.Username, user.Password, user.Auth, user.InfoId, user.Available)
	} else {
		ret, err = database.Db.Exec(`insert into user (username,password,auth,info_id,available) values(?,?,?,?,?)`, user.Username, user.Password, user.Auth, user.InfoId, user.Available)
	}
	var id int64
	if err != nil {
		return user, err
	}
	id, err = ret.LastInsertId()
	if err != nil {
		return user, err
	}
	user.Id = id
	return user, err
}

func UpdateUser(user User, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`update user set username=?, password=?, auth=?, info_id=?,available=? where id=?`, user.Username, user.Password, user.Auth, user.InfoId, user.Id, user.Available)
	} else {
		_, err = database.Db.Exec(`update user set username=?, password=?, auth=?, info_id=?,available=? where id=?`, user.Username, user.Password, user.Auth, user.InfoId, user.Id, user.Available)
	}
	return err
}

func UpdatePassword(uid int64,password string, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`update user set password=? where id=?`, password, uid)
	} else {
		_, err = database.Db.Exec(`update user set password=? where id=?`, password, uid)
	}
	return err
}

func UpdateUserStatus(uid int64, status int64, tx *sql.Tx) error{
	var err error
	if tx != nil {
		_, err = tx.Exec(`update user set available=? where id=?`, status, uid)
	} else {
		_, err = database.Db.Exec(`update user set available=? where id=?`, status, uid)
	}
	return err
}
func UpdateUserTutor(uid int64, tx *sql.Tx) error{
	var err error
	if tx != nil {
		_, err = tx.Exec(`update user set auth=? where id=?`, "tutor", uid)
	} else {
		_, err = database.Db.Exec(`update user set auth=? where id=?`, "tutor", uid)
	}
	return err
}
func UpdateUserManager(uid int64, tx *sql.Tx) error{
	var err error
	if tx != nil {
		_, err = tx.Exec(`update user set auth=? where id=?`, "manager", uid)
	} else {
		_, err = database.Db.Exec(`update user set auth=? where id=?`, "manager", uid)
	}
	return err
}
func LookupUserIdBySid(sid int64,tx *sql.Tx) (int64,error) {
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow(`select id from user where info_id = ? and auth=?;`, sid,"student")
	} else {
		rows = database.Db.QueryRow(`select id from user where info_id = ? and auth=?;`, sid,"student")
	}
	var id int64
	err := rows.Scan(&id)
	return id,err
}
func LookupUserIdByTid(tid int64,tx *sql.Tx) (int64,error) {
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow(`select id from user where info_id = ? and (auth=? or auth=?);`, tid,"tutor","manager")
	} else {
		rows = database.Db.QueryRow(`select id from user where info_id = ? and (auth=? or auth=?);`, tid,"tutor","manager")
	}
	var id int64
	err := rows.Scan(&id)
	return id,err
}