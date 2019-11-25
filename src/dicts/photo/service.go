package photo

import (
	"2019_2_Shtoby_shto/file_service/file"
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/models"
	transport "2019_2_Shtoby_shto/src/handle"
	"bufio"
	"bytes"
	"context"
	"os"
	"path"
	"time"
)

type HandlerPhotoService interface {
	DownloadPhoto(photo *bufio.Reader) (*models.Photo, error)
	GetPhotoByUser(photoID customType.StringUUID) ([]byte, error)
}

type service struct {
	transport.HandlerImpl
	cfg *config.Config
	db  database.IDataManager
	fl  file.IFileLoaderManagerClient
}

func CreateInstance(db database.IDataManager, cfg *config.Config, fileLoader file.IFileLoaderManagerClient) HandlerPhotoService {
	return &service{
		cfg: cfg,
		db:  db,
		fl:  fileLoader,
	}
}

func (s service) DownloadPhoto(photo *bufio.Reader) (*models.Photo, error) {
	photoPath := s.cfg.ImagePath
	if err := os.MkdirAll(photoPath, os.ModePerm); err != nil {
		return nil, err
	}
	newPhoto := &models.Photo{
		TimeLoad: time.Now(),
		Path:     photoPath,
	}
	if err := s.db.CreateRecord(newPhoto); err != nil {
		return nil, err
	}
	// внутрисервисное хранение - в последствии выпилится
	photoFile, err := os.Create(path.Join(photoPath, newPhoto.ID.String()+".jpg"))
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(photo); err != nil {
		return nil, err
	}
	if _, err := bufio.NewWriter(photoFile).Write(buf.Bytes()); err != nil {
		return nil, err
	}
	newFile := &file.File{
		ID:   newPhoto.ID.String(),
		Data: buf.Bytes(),
	}
	_, err = s.fl.DownloadFile(context.Background(), newFile)
	if err != nil {
		return nil, err
	}
	return newPhoto, nil
}

func (s service) GetPhotoByUser(photoID customType.StringUUID) ([]byte, error) {
	fileID := &file.FileID{
		ID: photoID.String(),
	}
	uploadFile, err := s.fl.UploadFile(context.Background(), fileID)
	if err != nil {
		return nil, err
	}
	return uploadFile.Data, nil
}
