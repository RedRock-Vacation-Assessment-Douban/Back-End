package dao

import "douban/model"

func SelectVideo1() ([]model.Video1, error) {
	var videos []model.Video1
	rows, err := dB.Query("SELECT id, title, context, url FROM video WHERE Id BETWEEN 1 AND 9")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var video model.Video1

		err = rows.Scan(&video.Id, &video.Title, &video.URL, &video.Context)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil
}

func SelectVideo2(id int) ([]model.Video2, error) {
	var videos []model.Video2
	rows, err := dB.Query("SELECT id, VideoSRC, URLInfo, Length FROM video WHERE Id = ?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var video model.Video2

		err = rows.Scan(&video.Id, &video.VideoSRC, &video.URLInfo, &video.Length)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil
}
