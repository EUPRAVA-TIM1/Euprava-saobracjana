package service

type SudService interface {
}

type sudServiceImpl struct {
	serviceUrl string
}

func NewSudService(url string) SudService {
	return sudServiceImpl{serviceUrl: url}
}
