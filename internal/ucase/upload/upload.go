// Package upload
// Automatic generated
package upload

import (
	"fmt"

	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/internal/repositories"
	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/tracer"

	ucase "cake-store/cake-store/internal/ucase/contract"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type upload struct {
	repo repositories.Uploader
}

// NewUpload new instance
func NewUpload(repo repositories.Uploader) ucase.UseCase {

	return &upload{repo: repo}
}

// Serve store upload data
func (u *upload) Serve(dctx *appctx.Data) appctx.Response {

	var (
		ctx = tracer.SpanStart(dctx.Request.Context(), "ucase.create")
		lf  = logger.NewFields(
			logger.EventName("Uploads"),
		)
	)

	file, header, _ := dctx.Request.FormFile("photo")

	fmt.Println(header.Filename)

	defer tracer.SpanFinish(ctx)
	var AccessKeyID = "N62BOTLXZ23X1TA0P35H"
	var SecretAccessKey = "c2H89J4eDYImF9OwAjdDQHQnDKwVwkKOmB5nKXOa"
	var MyRegion = "ap-southeast-3"

	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String("https://is3.cloudhost.id"),
		Region:      aws.String(MyRegion),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	})
	fmt.Println(err)
	svc := s3.New(sess)

	// Create a new PutObjectInput.
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String("lars-storage"),
		ACL:    aws.String("public-read"),
		Key:    aws.String(header.Filename),
		Body:   file,
	}

	// Upload the file to S3.
	_, err = svc.PutObject(putObjectInput)
	if err != nil {
		logger.WarnWithContext(ctx, fmt.Sprintf("validation error %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithError(err)
	}
	var filepath = "https://" + "lars-storage" + "." + "is3.cloudhost.id/" + header.Filename
	fmt.Println("File successfully uploaded to S3! :", filepath)

	logger.InfoWithContext(ctx, fmt.Sprintf("success store data to database"), lf...)
	return *appctx.NewResponse().
		WithMsgKey(consts.RespSuccess)
}
