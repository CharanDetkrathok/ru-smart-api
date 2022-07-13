package mr30r

import "github.com/jmoiron/sqlx"

type (
	mr30RepoDB struct {
		oracle_db *sqlx.DB
	}

	Mr30Repo struct {
	  	// เพิ่ม โครงสร้าง databse ตรงนี้
	}

	Mr30RepoInterface interface {
		GetMr30(course_year, course_semester string) (*[]Mr30Repo, error)
	}
)

func NewMr30Repo(oracle_db *sqlx.DB) Mr30RepoInterface {
	return &mr30RepoDB{oracle_db: oracle_db}
}
