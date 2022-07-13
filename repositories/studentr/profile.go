package studentr

func (r *studentRepoDB) GetStudentProfile(studentCode string) (student *StudentProfileRepo, err error) {

	student_info := StudentProfileRepo{}
	query := "ใส่ query string ตรงนี้"


	err = r.oracle_db.Get(&student_info, query, studentCode)
	if err != nil {
		return nil, err
	}

	student = &student_info

	return student, nil
}
