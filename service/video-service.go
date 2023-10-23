package service

import "gilab.com/pragmaticreviews/golang-gin-poc/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
	Delete(video entity.Video) []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

func (service *videoService) Delete(video entity.Video) []entity.Video {
	indexToDelete := -1
	for i, v := range service.videos {
		if v.Title == video.Title {
			indexToDelete = i
			break
		}
	}
	if indexToDelete >= 0 {
		service.videos = append(service.videos[:indexToDelete], service.videos[indexToDelete+1:]...)
	}
	return service.videos
}
