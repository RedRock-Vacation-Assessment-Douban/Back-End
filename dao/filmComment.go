package dao

import (
	"douban/model"
	"fmt"
)

// InsertFilmComment 发布影评
func InsertFilmComment(filmComment model.FilmComment) error {
	_, err := dB.Exec("INSERT INTO filmComment(MovieId, Name, Context, PostTime, StarNum, MovieName, URL) "+"values(?, ?, ?, ?, ?, ?, ?);", filmComment.MovieId, filmComment.Name, filmComment.Context, filmComment.PostTime, filmComment.StarNum, filmComment.MovieName, filmComment.URL)
	return err
}

// SelectFilmCommentById 通过id来搜索影评
func SelectFilmCommentById(filmCommentId int) (model.FilmComment, error) {
	var filmComment model.FilmComment

	row := dB.QueryRow("SELECT id, MovieId, Name, Context, PostTime, CommentNum, StarNum, Likes, Down, MovieName, URL FROM filmComment WHERE id = ? ", filmCommentId)
	if row.Err() != nil {
		return filmComment, row.Err()
	}

	err := row.Scan(&filmComment.Id, &filmComment.MovieId, &filmComment.Name, &filmComment.Context, &filmComment.PostTime, &filmComment.CommentNum, &filmComment.StarNum, &filmComment.Likes, &filmComment.Down, &filmComment.MovieName, &filmComment.URL)
	if err != nil {
		return filmComment, err
	}

	return filmComment, nil
}

// SelectNameByFCId 通过id查找发布用户
func SelectNameByFCId(FilmCommentId int) (string, error) {
	var filmComment model.FilmComment

	row := dB.QueryRow("SELECT Name FROM filmComment WHERE id = ? ", FilmCommentId)
	if row.Err() != nil {
		return filmComment.Name, row.Err()
	}

	err := row.Scan(&filmComment.Name)
	if err != nil {
		return filmComment.Name, err
	}

	return filmComment.Name, nil
}

// SelectFilmComment 查找影评
func SelectFilmComment(movieId int) ([]model.FilmComment, error) {
	var filmComments []model.FilmComment

	db.Model(&model.FilmComment{}).Where("MovieId = ?", movieId).Find(&filmComments)
	return filmComments, nil
}

// SelectFilmCommentByUsername 查找影评
func SelectFilmCommentByUsername(name string) ([]model.Personal, error) {
	var filmComments []model.Personal
	rows, err := dB.Query("SELECT MovieId, Id, MovieName, Name, Context, PostTime, CommentNum, StarNum, Likes, URL FROM filmComment WHERE Name = ?", name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var filmComment model.Personal

		err = rows.Scan(&filmComment.MovieId, &filmComment.Id, &filmComment.MovieName, &filmComment.Name, &filmComment.Context, &filmComment.PostTime, &filmComment.CommentNum, &filmComment.StarNum, &filmComment.Likes, &filmComment.URL)
		if err != nil {
			return nil, err
		}

		filmComments = append(filmComments, filmComment)
	}

	return filmComments, nil
}

// DeleteFilmComment 删除影评
func DeleteFilmComment(filmCommentId int) error {
	sqlStr := `delete from filmComment where Id=?`
	_, err := dB.Exec(sqlStr, filmCommentId)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

// FilmCommentLikes 给话题点赞
func FilmCommentLikes(id int) error {
	sqlStr := `update filmComment set Likes=Likes+1 where id = ?`
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// FilmCommentDown 给话题点踩
func FilmCommentDown(id int) error {
	sqlStr := `update filmComment set Down=Down+1 where id = ?`
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectMPFC() ([]model.MostPopularFC, error) {
	var MPFCs []model.MostPopularFC
	rows, err := dB.Query("SELECT id, MovieId ,Name, Context, StarNum, URL, MovieName FROM filmComment WHERE Id in (2,14,25,36,47)")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var MPFC model.MostPopularFC

		err = rows.Scan(&MPFC.Id, &MPFC.MovieId, &MPFC.Name, &MPFC.Context, &MPFC.StarNum, &MPFC.URL, &MPFC.MovieName)
		if err != nil {
			return nil, err
		}

		MPFCs = append(MPFCs, MPFC)
	}

	return MPFCs, nil
}

// SelectMNById 通过id查找
func SelectMNById(id int) (string, error) {
	var filmComment model.FilmComment

	row := dB.QueryRow("SELECT MovieName FROM filmComment WHERE MovieId = ? ", id)
	if row.Err() != nil {
		return filmComment.MovieName, row.Err()
	}

	err := row.Scan(&filmComment.MovieName)
	if err != nil {
		return filmComment.MovieName, err
	}

	return filmComment.MovieName, nil
}

// SelectURLByMId 通过id查找
func SelectURLByMId(id int) (string, error) {
	var filmComment model.FilmComment

	row := dB.QueryRow("SELECT URL FROM filmComment WHERE MovieId = ? ", id)
	if row.Err() != nil {
		return filmComment.URL, row.Err()
	}

	err := row.Scan(&filmComment.URL)
	if err != nil {
		return filmComment.URL, err
	}

	return filmComment.URL, nil
}
