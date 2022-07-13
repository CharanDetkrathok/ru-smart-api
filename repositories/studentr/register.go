package studentr

func (r *studentRepoDB) GetRegister(studentCode, courseYear, courseSemester string) (*[]RegisterRepo, error) {

	register := []RegisterRepo{}
	query := "ใส่ query string ตรงนี้"
	
	err := r.oracle_db.Select(&register, query)
	if err != nil {
		return nil, err
	}

	return &register, nil
}
