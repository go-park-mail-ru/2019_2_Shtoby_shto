package photo

import (
	"2019_2_Shtoby_shto/src/custom_type"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/utils"
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type HandlerPhotoService interface {
	DownloadPhoto(photoPath string, photo *bufio.Reader) (custom_type.StringUUID, error)
}

type service struct {
	db *database.DataManager
}

func CreateInstance(db *database.DataManager) HandlerPhotoService {
	return service{
		db: db,
	}
}

func (s service) DownloadPhoto(photoPath string, photo *bufio.Reader) (custom_type.StringUUID, error) {
	if err := os.MkdirAll(photoPath, os.ModePerm); err != nil {
		return "", err
	}
	newPhoto := Photo{
		TimeLoad: time.Now(),
		Path:     photoPath,
	}
	id, err := utils.GenerateUUID()
	if err != nil {
		return "", err
	}
	newPhoto.ID = custom_type.StringUUID(id.String())
	file, err := os.Create(path.Join(photoPath, newPhoto.ID.String()+".jpg"))
	if err != nil {
		return "", err
	}
	defer file.Close()
	photoData, err := ioutil.ReadAll(photo)
	if err != nil {
		return "", err
	}
	if _, err := bufio.NewWriter(file).Write(photoData); err != nil {
		return "", err
	}
	if err := s.db.CreateRecord(newPhoto); err != nil {
		return "", err
	}
	return newPhoto.ID, nil
}

func (s service) GetPhotoByUser() error {

	return nil
}
