package studentr

import (
	"github.com/jmoiron/sqlx"
)

type (
	studentRepoDB struct {
		oracle_db *sqlx.DB
	}

	StudentProfileRepo struct {
		// เพิ่มโครงสร้าง database ตรงนี้
	}

	RegisterRepo struct {
		// เพิ่มโครงสร้าง database ตรงนี้
	}

	PrepareTokenRepo struct {
		// เพิ่มโครงสร้าง Response ตรงนี้
	}

	StudentRepoInterface interface {
		GetStudentProfile(studentCode string) (*StudentProfileRepo, error)
		GetRegister(studentCode, courseYear, courseSemester string) (*[]RegisterRepo, error)
		Authentication(studentCode string) (*PrepareTokenRepo, error)
	}
)

func NewStudentRepo(oracle_db *sqlx.DB) StudentRepoInterface {
	return &studentRepoDB{oracle_db: oracle_db}
}
