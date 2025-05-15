package file

import (
	"github.com/go-resty/resty/v2"
	"mime/multipart"
)

type Client struct {
	client *resty.Client
}

func NewFileClient() *Client {
	client := resty.New().
		SetBaseURL("http://localhost:8090").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	return &Client{
		client: client,
	}
}

func (f Client) Upload(folder string, header *multipart.FileHeader) (string, error) {
	res, err := f.client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Content-Disposition", "form-data").
		SetPathParam("folder", folder).
		SetFile("file", header.Filename).
		Post("/folders/{folder}/files")
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func (f Client) Read(folder, file string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (f Client) Delete(folder, file string) error {
	//TODO implement me
	panic("implement me")
}
