package services

import "go-blog/dao"

var DbDao = dao.GetDaoInstance()

type Service struct {
}

func New() (service *Service) {
	return &Service{}
}
