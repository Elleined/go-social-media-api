package utils

type FileHandler interface {
	ReadFile()
	UploadFile()
	DeleteFile()
}

type FileHandlerImpl struct {
}

func NewFileHandler() FileHandler {
	return &FileHandlerImpl{}
}

func (f *FileHandlerImpl) DeleteFile() {
	panic("unimplemented")
}

func (f *FileHandlerImpl) ReadFile() {
	panic("unimplemented")
}

func (f *FileHandlerImpl) UploadFile() {
	panic("unimplemented")
}
