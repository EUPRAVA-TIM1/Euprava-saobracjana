package service

import (
	"EuprvaSaobracajna/data"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type FilesService interface {
	SavePdf(pdf []byte) (data.FileDto, error)
}

type filesServiceImpl struct {
	serviceUrl string
}

func NewFilesService(url string) FilesService {
	return filesServiceImpl{serviceUrl: url}
}

func (f filesServiceImpl) SavePdf(pdf []byte) (file data.FileDto, err error) {
	u, err := url.Parse(f.serviceUrl + "/api/files")
	if err != nil {
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "nalog.pdf")
	if err != nil {
		return
	}
	if _, err = part.Write(pdf); err != nil {
		return
	}
	writer.Close()

	client := http.Client{}
	req, err := http.NewRequest("POST", u.String(), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := client.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode == http.StatusCreated {
		err = json.NewDecoder(response.Body).Decode(&file)
		return
	}
	err = errors.New("can not reach FilesService")
	return
}
