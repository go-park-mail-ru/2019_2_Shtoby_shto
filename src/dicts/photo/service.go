package photo

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/models"
	transport "2019_2_Shtoby_shto/src/handle"
	"bufio"
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
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
	svc *s3.S3
	cfg *config.Config
	db  database.IDataManager
}

func CreateInstance(db database.IDataManager, cfg *config.Config) HandlerPhotoService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(cfg.StorageRegion),
		Endpoint: aws.String(cfg.StorageEndpoint),
	}))
	svc := s3.New(sess)
	return &service{
		svc: svc,
		cfg: cfg,
		db:  db,
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
	file, err := os.Create(path.Join(photoPath, newPhoto.ID.String()+".jpg"))
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(photo); err != nil {
		return nil, err
	}
	if _, err := bufio.NewWriter(file).Write(buf.Bytes()); err != nil {
		return nil, err
	}
	var acl = s3.ObjectCannedACLPublicRead
	r := bytes.NewReader(buf.Bytes())
	_, err = s.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.cfg.StorageBucket),
		Key:    aws.String(newPhoto.ID.String()),
		ACL:    &acl,
		Body:   r,
	})
	if err != nil {
		return nil, err
	}
	return newPhoto, nil
}

func (s service) GetPhotoByUser(photoID customType.StringUUID) ([]byte, error) {
	out, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.cfg.StorageBucket),
		Key:    aws.String(photoID.String()),
	})
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(out.Body)
}
