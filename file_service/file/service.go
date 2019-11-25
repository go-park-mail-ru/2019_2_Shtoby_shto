package file

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
)

type IFileLoaderManager interface {
	DownloadFile(context.Context, *File) (*Nothing, error)
	UploadFile(context.Context, *FileID) (*File, error)
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

func (s service) DownloadFile(ctx context.Context, file *File) (*Nothing, error) {
	acl := s3.ObjectCannedACLPublicRead
	r := bytes.NewReader(file.Data)
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.storageBucket),
		Key:    aws.String(file.ID),
		ACL:    &acl,
		Body:   r,
	})
	if err != nil {
		return nil, err
	}
	return &Nothing{}, nil
}

func (s service) UploadFile(ctx context.Context, fileID *FileID) (*File, error) {
	out, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.storageBucket),
		Key:    aws.String(fileID.ID),
	})
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(out.Body)
	file := &File{
		ID:   fileID.ID,
		Data: data,
	}
	return file, nil
}
