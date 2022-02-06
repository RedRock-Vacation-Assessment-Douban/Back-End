package dao

import "douban/model"

// SelectCelebrityById 通过id来搜索影人
func SelectCelebrityById(Id int) (model.Celebrity, error) {
	var celebrity model.Celebrity

	row := dB.QueryRow("SELECT id, name, info, synopsis, url, award, urlinfo1, urlinfo2, urlinfo3, nameinfo1, nameinfo2, nameinfo3, yearinfo FROM celebrity WHERE id = ? ", Id)
	if row.Err() != nil {
		return celebrity, row.Err()
	}

	err := row.Scan(&celebrity.Id, &celebrity.Name, &celebrity.Info, &celebrity.Synopsis, &celebrity.URL, &celebrity.Award, &celebrity.URLInfo1, &celebrity.URLInfo2, &celebrity.URLInfo3, &celebrity.NameInfo1, &celebrity.NameInfo2, &celebrity.NameInfo3, &celebrity.YearInfo)
	if err != nil {
		return celebrity, err
	}

	return celebrity, nil
}

func Search2(context string) ([]model.Search2, error) {
	var celebrities []model.Search2

	rows, err := dB.Query("SELECT name, url FROM celebrity WHERE Name LIKE ?", context)

	defer rows.Close()
	for rows.Next() {
		var celebrity model.Search2

		err = rows.Scan(&celebrity.Name, &celebrity.URL)
		if err != nil {
			return nil, err
		}

		celebrities = append(celebrities, celebrity)
	}

	return celebrities, nil
}
