package services

import "go-blog/models"

func (s *Service) GetAllPosts() (posts []models.Post, err error) {
	return DbDao.SelectAllPost()
}
