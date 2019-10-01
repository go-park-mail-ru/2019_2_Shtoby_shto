package photo

import "os"

type HandlerPhotoService interface {
	DownloadPhoto(photoPath string, photo []byte) error
}

type service struct {
}

func (s service) NewPhotoService() HandlerPhotoService {
	return service{}
}

func (s service) DownloadPhoto(photoPath string, photo []byte) error {
	if err := os.MkdirAll(photoPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}
