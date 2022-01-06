package topic

import (
	"backend/database/internel"
	"backend/database/tutor"
	"database/sql"
)

type Topic struct {
	Id                 int64  `form:"id"`
	Name               string `form:"name"`
	Type               string `form:"type"`
	Source             string `form:"source"`
	Tutor              tutor.Tutor
	Profile            string `form:"profile"`
	MajorRequirement   string `form:"major_requirement"`
	StudentRequirement string `form:"student_requirement"`
	Passed             int8
	Published          int8
}

func GetTopicById(tid int64, tx *sql.Tx) (Topic, error) {
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id WHERE topic.id =  ?", tid)
	} else {
		rows = database.Db.QueryRow("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id WHERE topic.id = ?", tid)
	}
	var topic Topic
	err := rows.Scan(
		&topic.Id,
		&topic.Name,
		&topic.Type,
		&topic.Source,
		&topic.Tutor.Id,
		&topic.Profile,
		&topic.MajorRequirement,
		&topic.StudentRequirement,
		&topic.Passed,
		&topic.Published,
		&topic.Tutor.Id,
		&topic.Tutor.Name,
		&topic.Tutor.Sex,
		&topic.Tutor.Birthday,
		&topic.Tutor.EducationalBackground,
		&topic.Tutor.Title,
		&topic.Tutor.ResearchDirection,
		&topic.Tutor.PhoneNumber,
		&topic.Tutor.Email,
	)
	return topic, err
}
func GetTopicList(tx *sql.Tx) ([]Topic, error) {
	topics := make([]Topic, 0)
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.Query("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id ")

	} else {
		rows, err = database.Db.Query("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id ")
	}
	if err != nil {
		return topics, err
	}
	var topic Topic
	for rows.Next() {
		err = rows.Scan(
			&topic.Id,
			&topic.Name,
			&topic.Type,
			&topic.Source,
			&topic.Tutor.Id,
			&topic.Profile,
			&topic.MajorRequirement,
			&topic.StudentRequirement,
			&topic.Passed,
			&topic.Published,
			&topic.Tutor.Id,
			&topic.Tutor.Name,
			&topic.Tutor.Sex,
			&topic.Tutor.Birthday,
			&topic.Tutor.EducationalBackground,
			&topic.Tutor.Title,
			&topic.Tutor.ResearchDirection,
			&topic.Tutor.PhoneNumber,
			&topic.Tutor.Email,
		)
		topics = append(topics, topic)
	}
	return topics, err
}
func GetTopicListByTutorId(tid int64, tx *sql.Tx) ([]Topic, error) {
	topics := make([]Topic, 0)
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.Query("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id where tutor.id = ?", tid)

	} else {
		rows, err = database.Db.Query("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id where tutor.id = ?", tid)
	}
	if err != nil {
		return topics, err
	}
	var topic Topic
	for rows.Next() {
		err = rows.Scan(
			&topic.Id,
			&topic.Name,
			&topic.Type,
			&topic.Source,
			&topic.Tutor.Id,
			&topic.Profile,
			&topic.MajorRequirement,
			&topic.StudentRequirement,
			&topic.Passed,
			&topic.Published,
			&topic.Tutor.Id,
			&topic.Tutor.Name,
			&topic.Tutor.Sex,
			&topic.Tutor.Birthday,
			&topic.Tutor.EducationalBackground,
			&topic.Tutor.Title,
			&topic.Tutor.ResearchDirection,
			&topic.Tutor.PhoneNumber,
			&topic.Tutor.Email,
		)
		topics = append(topics, topic)
	}
	return topics, err
}

func CreateTopic(topic Topic, tx *sql.Tx) (Topic, error) {
	var ret sql.Result
	var err error
	if tx != nil {
		ret, err = tx.Exec(`
		insert into topic (name,type,source,tutor,profile,major_requirement,student_requirement) values(?,?,?,?,?,?,?)
		`, topic.Name, topic.Type, topic.Source, topic.Tutor.Id, topic.Profile, topic.MajorRequirement, topic.StudentRequirement)
	} else {
		ret, err = database.Db.Exec(`
		insert into topic (name,type,source,tutor,profile,major_requirement,student_requirement) values(?,?,?,?,?,?,?)
		`, topic.Name, topic.Type, topic.Source, topic.Tutor.Id, topic.Profile, topic.MajorRequirement, topic.StudentRequirement)
	}
	var id int64
	if err != nil {
		return topic, err
	}
	id, err = ret.LastInsertId()
	if err != nil {
		return topic, err
	}
	topic.Id = id
	return topic, err
}

func UpdateTopic(topic Topic, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`
		update topic set name=?, type=?, source=?, tutor=?, profile=?, major_requirement=?, student_requirement=? where id=?
		`,
			topic.Name, topic.Type, topic.Source, topic.Tutor.Id, topic.Profile, topic.MajorRequirement, topic.StudentRequirement, topic.Id)
	} else {
		_, err = database.Db.Exec(`
		update topic set name=?, type=?, source=?, tutor=?, profile=?, major_requirement=?, student_requirement=? where id=?
		`,
			topic.Name, topic.Type, topic.Source, topic.Tutor.Id, topic.Profile, topic.MajorRequirement, topic.StudentRequirement, topic.Id)
	}
	return err
}

func SetPassedStatus(topic Topic, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`
		update topic set passed=? where id=?
		`,
			topic.Passed, topic.Id)
	} else {
		_, err = database.Db.Exec(`
		update topic set passed=? where id=?
		`,
			topic.Passed, topic.Id)
	}
	return err
}

func SetPublishedStatus(topic Topic, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`
		update topic set published=? where id=?
		`,
			topic.Published, topic.Id)
	} else {
		_, err = database.Db.Exec(`
		update topic set published=? where id=?
		`,
			topic.Published, topic.Id)
	}
	return err
}

func GetTopicListByPublishStatus(status int8, tx *sql.Tx) ([]Topic, error) {
	topics := make([]Topic, 0)
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.Query("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id where topic.published = ?", status)

	} else {
		rows, err = database.Db.Query("SELECT * FROM `topic` INNER JOIN `tutor` on topic.tutor = tutor.id where topic.published = ?", status)
	}
	if err != nil {
		return topics, err
	}
	var topic Topic
	for rows.Next() {
		err = rows.Scan(
			&topic.Id,
			&topic.Name,
			&topic.Type,
			&topic.Source,
			&topic.Tutor.Id,
			&topic.Profile,
			&topic.MajorRequirement,
			&topic.StudentRequirement,
			&topic.Passed,
			&topic.Published,
			&topic.Tutor.Id,
			&topic.Tutor.Name,
			&topic.Tutor.Sex,
			&topic.Tutor.Birthday,
			&topic.Tutor.EducationalBackground,
			&topic.Tutor.Title,
			&topic.Tutor.ResearchDirection,
			&topic.Tutor.PhoneNumber,
			&topic.Tutor.Email,
		)
		topics = append(topics, topic)
	}
	return topics, err
}