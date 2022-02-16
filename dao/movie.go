package dao

import "douban/model"

// SelectMovieById 通过id来搜索电影
func SelectMovieById(movieId int) (model.Movie, error) {
	var movie model.Movie

	row := dB.QueryRow("SELECT id, name, year, director, screenwriter, starring, type, country, language, length, imdb, starnum, score, star, havewatched, wanttowatch, synopsis, URL, movieURL, peopleURL, NameInfo, CoverInfo FROM movie WHERE id = ? ", movieId)
	if row.Err() != nil {
		return movie, row.Err()
	}

	err := row.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Type, &movie.Country, &movie.Language, &movie.Length, &movie.IMDb, &movie.StarNum, &movie.Score, &movie.Star, &movie.HaveWatched, &movie.WantToWatch, &movie.Synopsis, &movie.URL, &movie.MovieURL, &movie.PeopleURL, &movie.NameInfo, &movie.CoverInfo)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

// SelectMovie1 查找主页1
func SelectMovie1() ([]model.MovieBrief1, error) {
	var movies []model.MovieBrief1
	rows, err := dB.Query("SELECT id, name, URL, Score FROM movie WHERE Id BETWEEN 1 AND 35")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie model.MovieBrief1

		err = rows.Scan(&movie.Id, &movie.Name, &movie.URL, &movie.Score)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectMovie2 查找主页2
func SelectMovie2() ([]model.MovieBrief, error) {
	var movies []model.MovieBrief
	rows, err := dB.Query("SELECT id, name, URL FROM movie WHERE Id BETWEEN 36 AND 85")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie model.MovieBrief

		err = rows.Scan(&movie.Id, &movie.Name, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectMovie3 查找主页3
func SelectMovie3() ([]model.MovieBrief, error) {
	var movies []model.MovieBrief
	rows, err := dB.Query("SELECT id, name, URL FROM movie WHERE Id BETWEEN 86 AND 135")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie model.MovieBrief

		err = rows.Scan(&movie.Id, &movie.Name, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectURLById 通过id查找URL
func SelectURLById(movieId int) string {
	var movie model.Movie

	row := dB.QueryRow("SELECT URL FROM movie WHERE id = ? ", movieId)
	if row.Err() != nil {
		return movie.URL
	}

	err := row.Scan(&movie.URL)
	if err != nil {
		return movie.URL
	}

	return movie.URL
}

// SelectMovie 搜索电影
func SelectMovie() ([]model.Rank1, error) {
	var movies []model.Rank1

	rows, err := dB.Query("SELECT id, name, year, starring, country, starnum, score, havewatched, URL FROM movie WHERE id BETWEEN 1 AND 10")

	defer rows.Close()
	for rows.Next() {
		var movie model.Rank1

		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Starring, &movie.Country, &movie.StarNum, &movie.Score, &movie.HaveWatched, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectMovieRank1 搜索电影
func SelectMovieRank1() ([]model.Rank2, error) {
	var movies []model.Rank2

	rows, err := dB.Query("SELECT id, name FROM movie WHERE id BETWEEN 1 AND 10")

	defer rows.Close()
	for rows.Next() {
		var movie model.Rank2

		err = rows.Scan(&movie.Id, &movie.Name)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectMovieRank250 搜索电影
func SelectMovieRank250() ([]model.Rank250, error) {
	var movies []model.Rank250

	rows, err := dB.Query("SELECT name, URL FROM movie WHERE id BETWEEN 136 AND 147")

	defer rows.Close()
	for rows.Next() {
		var movie model.Rank250

		err = rows.Scan(&movie.Name, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func Search(context string) ([]model.Search, error) {
	var movies []model.Search

	rows, err := dB.Query("SELECT id, name, year, director, starring, type, country, length, starnum, score, havewatched, url FROM movie WHERE Name LIKE ?", context)

	defer rows.Close()
	for rows.Next() {
		var movie model.Search

		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Starring, &movie.Type, &movie.Country, &movie.Length, &movie.StarNum, &movie.Score, &movie.HaveWatched, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// SelectMovieUSA 搜索电影
func SelectMovieUSA() ([]model.USA, error) {
	var movies []model.USA

	rows, err := dB.Query("SELECT id, name, BoxOffice FROM USA WHERE id BETWEEN 1 AND 10")

	defer rows.Close()
	for rows.Next() {
		var movie model.USA

		err = rows.Scan(&movie.Id, &movie.Name, &movie.BoxOffice)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func Classify(mold string, country string) ([]model.Classify, error) {
	var movies []model.Classify
	//union select id, name, score, url from movie where Country LIKE ?
	rows, err := dB.Query("SELECT id, name, score, url FROM movie WHERE Type LIKE ? and Country LIKE ?", mold, country)

	defer rows.Close()
	for rows.Next() {
		var movie model.Classify
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Score, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func All() ([]model.Classify, error) {
	var movies []model.Classify
	rows, err := dB.Query("SELECT id, name, score, url FROM movie WHERE Type !='' ")

	defer rows.Close()
	for rows.Next() {
		var movie model.Classify
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Score, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func Classify1(country string) ([]model.Classify, error) {
	var movies []model.Classify
	rows, err := dB.Query("SELECT id, name, score, url FROM movie WHERE Country LIKE ?", country)

	defer rows.Close()
	for rows.Next() {
		var movie model.Classify
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Score, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func Classify2(mold string) ([]model.Classify, error) {
	var movies []model.Classify
	rows, err := dB.Query("SELECT id, name, score, url FROM movie WHERE Type LIKE ?", mold)

	defer rows.Close()
	for rows.Next() {
		var movie model.Classify
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Score, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func ClassifyRank(mold string) ([]model.ClassifyRank, error) {
	var movies []model.ClassifyRank
	rows, err := dB.Query("SELECT id, name, year, starring, type, country, starnum, score, havewatched, url FROM movie WHERE Type LIKE ? ORDER BY Score desc", mold)

	defer rows.Close()
	for rows.Next() {
		var movie model.ClassifyRank
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Starring, &movie.Type, &movie.Country, &movie.StarNum, &movie.Score, &movie.HaveWatched, &movie.URL)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func SelectMovieNameById(movieId int) string {
	var movie model.ShortComment

	row := dB.QueryRow("SELECT MovieName FROM shortComment WHERE MovieId = ? ", movieId)
	if row.Err() != nil {
		return movie.MovieName
	}

	err := row.Scan(&movie.MovieName)
	if err != nil {
		return movie.MovieName
	}

	return movie.MovieName
}
