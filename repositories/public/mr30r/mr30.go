package mr30r

func (r *mr30RepoDB) GetMr30(course_year, course_semester string) (*[]Mr30Repo, error) {

	mr30_info := []Mr30Repo{}
	query := "ใส่ query string ตรงนี้"

	err := r.oracle_db.Select(&mr30_info, query)

	if err != nil {
		return nil, err
	}

	return &mr30_info, nil
}
