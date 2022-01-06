package tutor

import (
	"backend/database/internel"
	"database/sql"
)

type Tutor struct {
	Id                    int64  `form:"id"`
	Name                  string `form:"name"`
	Sex                   string `form:"sex"`
	Birthday              string `form:"birthday"`
	EducationalBackground string `form:"EducationalBackground"`
	Title                 string `form:"title"`
	ResearchDirection     string `form:"ResearchDirection"`
	PhoneNumber           string `form:"PhoneNumber"`
	Email                 string `form:"Email"`
}

func GetTutorById(tid int64, tx *sql.Tx) (Tutor, error) {
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow("select * from tutor where id = ?", tid)
	} else {
		rows = database.Db.QueryRow("select * from tutor where id = ?", tid)
	}
	var tutor Tutor
	err := rows.Scan(&tutor.Id, &tutor.Name, &tutor.Sex, &tutor.Birthday, &tutor.EducationalBackground, &tutor.Title, &tutor.ResearchDirection, &tutor.PhoneNumber, &tutor.Email)
	return tutor, err
}

func CreateTutor(tutor Tutor, tx *sql.Tx) (Tutor, error) {
	var ret sql.Result
	var err error
	if tx != nil {
		ret, err = tx.Exec(`
		insert into tutor (name,sex,birthday,educational_background,title,research_direction,phone_number,email) values(?,?,?,?,?,?,?,?)
		`, tutor.Name, tutor.Sex, tutor.Birthday, tutor.EducationalBackground, tutor.Title, tutor.ResearchDirection, tutor.PhoneNumber, tutor.Email)
	} else {
		ret, err = database.Db.Exec(
			`
				insert into tutor (name,sex,birthday,educational_background,title,research_direction,phone_number,email) values(?,?,?,?,?,?,?,?)
			`, tutor.Name, tutor.Sex, tutor.Birthday, tutor.EducationalBackground, tutor.Title, tutor.ResearchDirection, tutor.PhoneNumber, tutor.Email)
	}
	var id int64
	if err != nil {
		return tutor, err
	}
	id, err = ret.LastInsertId()
	if err != nil {
		return tutor, err
	}
	tutor.Id = id
	return tutor, err
}

func UpdateTutor(tutor Tutor, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`
		update tutor set name=?, sex=?, birthday=?, educational_background=?, title=?, research_direction=?, phone_number=?, email=? where id=?
		`,
			tutor.Name, tutor.Sex, tutor.Birthday, tutor.EducationalBackground, tutor.Title, tutor.ResearchDirection, tutor.PhoneNumber, tutor.Email, tutor.Id)
	} else {
		_, err = database.Db.Exec(`
		update tutor set name=?, sex=?, birthday=?, educational_background=?, title=?, research_direction=?, phone_number=?, email=? where id=?
		`,
			tutor.Name, tutor.Sex, tutor.Birthday, tutor.EducationalBackground, tutor.Title, tutor.ResearchDirection, tutor.PhoneNumber, tutor.Email, tutor.Id)
	}
	return err
}
