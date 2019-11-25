package fileLoader

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type IFileLoaderManager interface {
	DownloadFile(id string, uploadFile []byte) error
	UploadFile(id string) (io.ReadCloser, error)
}

type service struct {
	svc           *s3.S3
	storageBucket string
}

func CreateFileLoaderInstance(storageRegion, storageEndpoint, storageBucket string) IFileLoaderManager {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(storageRegion),
		Endpoint: aws.String(storageEndpoint),
	}))
	svc := s3.New(sess)
	return service{
		svc:           svc,
		storageBucket: storageBucket,
	}
}

func (s service) DownloadFile(id string, uploadFile []byte) error {
	acl := s3.ObjectCannedACLPublicRead
	r := bytes.NewReader(uploadFile)
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.storageBucket),
		Key:    aws.String(id),
		ACL:    &acl,
		Body:   r,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s service) UploadFile(id string) (io.ReadCloser, error) {
	out, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.storageBucket),
		Key:    aws.String(id),
	})
	if err != nil {
		return nil, err
	}
	return out.Body, nil
}
