package studentr

func (r *studentRepoDB) Authentication(studentCode string) (token *PrepareTokenRepo, err error) {

	tempToken := PrepareTokenRepo{}
	query := "ใส่ query string ตรงนี้"


	err = r.oracle_db.Get(&tempToken, query, studentCode)
	if err != nil {
		return nil, err
	}

	token = &tempToken

	return token, nil
}
