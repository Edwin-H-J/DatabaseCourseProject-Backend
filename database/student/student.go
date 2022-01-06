package student

import (
	"backend/database/internel"
	"database/sql"
)

type Student struct{
	Id		int64		`form:"id"`
	Name	string		`form:"name"`
	Sex		string		`form:"sex"`
	Major	string		`form:"major"`
	Class	string		`form:"class"`
	PhoneNumber	string	`form:"phoneNumber"`
	Email	string		`form:"Email"`
	Remark	string		`form:"remark"`
}

func GetStudentById(tid int64,tx *sql.Tx) (Student,error){
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow("select * from student where id = ?", tid)
	}else{
		rows = database.Db.QueryRow("select * from student where id = ?", tid)
	}
	var student Student
	err := rows.Scan(&student.Id, &student.Name, &student.Sex, &student.Major, &student.Class,&student.PhoneNumber,&student.Email,&student.Remark)
	return student,err
}

func CreateStudent(student Student,tx *sql.Tx) (Student,error){
	var ret sql.Result
	var err error
	if tx != nil {
		ret,err = tx.Exec(`
		insert into student (name,sex,major,class,phone_number,email,remark) values(?,?,?,?,?,?,?)
		`,student.Name,student.Sex,student.Major,student.Class,student.PhoneNumber,student.Email,student.Remark)
	}else{
		ret,err = database.Db.Exec(
			`
			insert into student (name,sex,major,class,phone_number,email,remark) values(?,?,?,?,?,?,?)
		`,student.Name,student.Sex,student.Major,student.Class,student.PhoneNumber,student.Email,student.Remark)
	}
	var id int64
	if err != nil {
		return student,err
	}
	id,err = ret.LastInsertId()
	if err != nil {
		return student,err
	}
	student.Id = id
	return student,err
}

func UpdateStudent(student Student,tx *sql.Tx) error{
	var err error
	if tx != nil {
		_,err = tx.Exec(`
		update student set name=?, sex=?, major=?, class=?, phone_number=?, email=?, remark=? where id=?
		`,
		student.Name,student.Sex,student.Major,student.Class,student.PhoneNumber,student.Email,student.Remark,student.Id)
	}else{
		_,err = database.Db.Exec(`
		update student set name=?, sex=?, major=?, class=?, phone_number=?, email=?, remark=? where id=?
		`,
		student.Name,student.Sex,student.Major,student.Class,student.PhoneNumber,student.Email,student.Remark,student.Id)
	}
	return err
}


