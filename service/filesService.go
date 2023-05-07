package service

type FilesService interface {
}

type filesServiceImpl struct {
	serviceUrl string
}

func NewFilesService(url string) FilesService {
	return filesServiceImpl{serviceUrl: url}
}
