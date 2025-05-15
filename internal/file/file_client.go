package file

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"mime/multipart"
	"os"
)

type Client struct {
	client *resty.Client
}

func NewFileClient(client *resty.Client) *Client {
	return &Client{
		client: client,
	}
}

func (f Client) Upload(folder string, uploaded multipart.File, header *multipart.FileHeader) (string, error) {
	// Create a temporary file to store the uploaded content
	tempFile, err := os.CreateTemp("", "upload-*-"+header.Filename)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			return
		}
	}(tempFile.Name()) // Clean up after upload

	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			return
		}
	}(tempFile)

	// Copy uploaded content to the temp file
	_, err = io.Copy(tempFile, uploaded)
	if err != nil {
		return "", fmt.Errorf("failed to copy uploaded content: %w", err)
	}

	res, err := f.client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Content-Disposition", "form-data").
		SetPathParam("folder", folder).
		SetFile("file", tempFile.Name()).
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
