package process

import (
	"backend/database/internel"
	"database/sql"
)

type Process struct {
	Id                  int64  `json:"id" form:"id"`
	ProcessStatus       int64  `json:"process_status" form:"process_status"`
	TaskBook            string `json:"task_book" form:"task_book"`
	LiteratureReview    string `json:"literature_review" form:"literature_review"`
	Proposal            string `json:"proposal" form:"proposal"`
	DocumentTranslation string `json:"document_translation" form:"document_translation"`
	MidTermReport       string `json:"mid_term_report" form:"mid_term_report"`
	MidTermResult       string `json:"mid_term_result" form:"mid_term_result"`
	Paper               string `json:"paper" form:"paper"`
	TutorReview         string `json:"tutor_review" form:"tutor_review"`
	PeerReview          string `json:"peer_review" form:"peer_review"`
	DefendResult        string `json:"defend_result" form:"defend_result"`
}

func CreateProcess(tx *sql.Tx) (int64, error) {
	var ret sql.Result
	var err error
	if tx != nil {
		ret, err = tx.Exec(`insert into process (process_status,task_book,literature_review,proposal,document_translation,mid_term_report,mid_term_result,paper,tutor_review,peer_review,defend_result) values(?,?,?,?,?,?,?,?,?,?,?)`, 0,"","","","","","","","","","")
	} else {
		ret, err = database.Db.Exec(`insert into process (process_status,task_book,literature_review,proposal,document_translation,mid_term_report,mid_term_result,paper,tutor_review,peer_review,defend_result) values(?,?,?,?,?,?,?,?,?,?,?)`, 0,"","","","","","","","","","")
	}
	if err != nil {
		return 0, err
	}
	return ret.LastInsertId()
}
func GetProcessByPid(pid int64, tx *sql.Tx)(Process,error){
	var err error
	var rows *sql.Row
	if tx != nil {
		rows= tx.QueryRow(`select * from process where id = ?`,pid)
	} else {
		rows = database.Db.QueryRow(`select * from process where id = ?`,pid)
	}
	var process Process
	err = rows.Scan(
		&process.Id,
		&process.ProcessStatus,
		&process.TaskBook,
		&process.LiteratureReview,
		&process.Proposal,
		&process.DocumentTranslation,
		&process.MidTermReport,
		&process.MidTermResult,
		&process.Paper,
		&process.TutorReview,
		&process.PeerReview,
		&process.DefendResult)
	return process,err
}
func UpdateStatus(status int,pid int64, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`update process set process_status=? where id=?`, status,pid)
	} else {
		_, err = database.Db.Exec(`update process set process_status=? where id=?`, status,pid)
	}
	return err
}
func UpdateProp(prop string, detail string,pid int64, tx *sql.Tx) error {
	var err error
	if tx != nil {
		_, err = tx.Exec(`update process set `+prop+`=? where id=?` , detail,pid)
	} else {
		_, err = database.Db.Exec(`update process set `+prop+`=? where id=?` , detail,pid)
	}
	return err
}
func LookUpSidAndTidByPid(pid int64, tx *sql.Tx) (int64,int64,error) {
	var err error
	var rows *sql.Row
	if tx != nil {
		rows = tx.QueryRow(
`SELECT
	tutor.id AS tid,
	student.id AS sid 
FROM
	process
	INNER JOIN selected ON process.id = selected.process_id
	INNER JOIN topic ON topic.id = selected.tid
	INNER JOIN student ON student.id = selected.sid
	INNER JOIN tutor ON tutor.id = topic.tutor
WHERE
	process.id = ?`, pid)
	} else {
		rows= database.Db.QueryRow(
`SELECT
	tutor.id AS tid,
	student.id AS sid 
FROM
	process
	INNER JOIN selected ON process.id = selected.process_id
	INNER JOIN topic ON topic.id = selected.tid
	INNER JOIN student ON student.id = selected.sid
	INNER JOIN tutor ON tutor.id = topic.tutor
WHERE
	process.id = ?`, pid)
	}
	var tid,sid int64
	err = rows.Scan(&tid,&sid)
	return tid,sid,err
}
