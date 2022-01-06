package selected
import (
	"backend/database/internel"
	"database/sql"
)
type Selected struct{
	Tid int64
	Sid int64
	TutorCheck int8
	ManagerCheck int8
	Published int8
	ProcessId interface{}
	Topic interface{}
	Student interface{}
}


func CreateSelected(Tid int64,Sid int64,tx *sql.Tx) (error){
	var err error
	if tx != nil {
		_,err = tx.Exec(`insert into selected (tid,sid) values(?,?)`,Tid,Sid)
	}else{
		_,err = database.Db.Exec(`insert into selected (tid,sid) values(?,?)`,Tid,Sid)
	}
	return err
}
func GetSelectedBySid(Sid int64,tx *sql.Tx)(Selected,error){
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow("select * from selected where sid = ?", Sid)
	}else{
		rows = database.Db.QueryRow("select * from selected where sid = ?", Sid)
	}
	var selected Selected
	err := rows.Scan(&selected.Tid,&selected.Sid,&selected.TutorCheck,&selected.ManagerCheck,&selected.Published,&selected.ProcessId)
	return selected,err
}
func SetTutorComfirm(Tid int64,Sid int64,status int8,tx *sql.Tx)(error){
	var err error
	if tx != nil {
		_,err = tx.Exec(`update selected set tutor_check=? where tid=? and sid=?`,status,Tid,Sid)
	}else{
		_,err = database.Db.Exec(`update selected set tutor_check=? where tid=? and sid=?`,status,Tid,Sid)
	}
	return err
	
}
func SetPublished(Tid int64,Sid int64,Pid int64,status int8,tx *sql.Tx)(error){
	var err error
	if tx != nil {
		_,err = tx.Exec(`update selected set published=?,process_id=? where tid=? and sid=?`,status,Pid,Tid,Sid)
	}else{
		_,err = database.Db.Exec(`update selected set published=?,process_id=? where tid=? and sid=?`,status,Pid,Tid,Sid)
	}
	return err
	
}
func SetManagerComfirm(Tid int64,Sid int64,status int8,tx *sql.Tx)(error){
	var err error
	if tx != nil {
		_,err = tx.Exec(`update selected set manage_check=? where tid=? and sid=?`,status,Tid,Sid)
	}else{
		_,err = database.Db.Exec(`update selected set manage_check=? where tid=? and sid=?`,status,Tid,Sid)
	}
	return err
	
}
func DeleteSelected(Tid int64,Sid int64,tx *sql.Tx) (error){
	var err error
	if tx != nil {
		_,err = tx.Exec(`delete from selected where tid=? and sid=?`,Tid,Sid)
	}else{
		_,err = database.Db.Exec(`delete from selected where tid=? and sid=?`,Tid,Sid)
	}
	return err
}
func GetTutorConfirmSelected(tx *sql.Tx) ([]Selected, error){
	selections := make([]Selected, 0)
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.Query("SELECT * FROM `selected` where tutor_check = ?", 1)

	} else {
		rows, err = database.Db.Query("SELECT * FROM `selected` where tutor_check = ?", 1)
	}
	if err != nil {
		return selections, err
	}
	var selected Selected
	for rows.Next() {
		err := rows.Scan(&selected.Tid,&selected.Sid,&selected.TutorCheck,&selected.ManagerCheck,&selected.Published,&selected.ProcessId)
		if err!=nil {
			return selections,err
		}
		selections = append(selections, selected)
	}
	return selections,err
}
func GetSelectedByTutorId(tutorId int64,tx *sql.Tx) ([]Selected, error){
	selections := make([]Selected, 0)
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.Query("SELECT * FROM selected WHERE tid in (SELECT id FROM topic WHERE tutor = ?)", tutorId)

	} else {
		rows, err = database.Db.Query("SELECT * FROM selected WHERE tid in (SELECT id FROM topic WHERE tutor = ?)", tutorId)
	}
	if err != nil {
		return selections, err
	}
	var selected Selected
	for rows.Next() {
		err := rows.Scan(&selected.Tid,&selected.Sid,&selected.TutorCheck,&selected.ManagerCheck,&selected.Published,&selected.ProcessId)
		if err!=nil {
			return selections,err
		}
		selections = append(selections, selected)
	}
	return selections,err
}

func GetSelectedByPid(Sid int64,tx *sql.Tx)(Selected,error){
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow("select * from selected where sid = ?", Sid)
	}else{
		rows = database.Db.QueryRow("select * from selected where sid = ?", Sid)
	}
	var selected Selected
	err := rows.Scan(&selected.Tid,&selected.Sid,&selected.TutorCheck,&selected.ManagerCheck,&selected.Published,&selected.ProcessId)
	return selected,err
}