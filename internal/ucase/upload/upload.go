// Package upload
// Automatic generated
package upload

import (
	"fmt"
	"io/ioutil"
	"mime"
	"path/filepath"

	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/internal/repositories"
	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/storage"
	"cake-store/cake-store/pkg/tracer"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	ucase "cake-store/cake-store/internal/ucase/contract"
)

var cfg = appctx.NewConfig()

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

	//fmt.Println(dctx.Request.Context())

	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String("https://is3.cloudhost.id"),
		Region:      aws.String(cfg.AWS.Region),
		Credentials: credentials.NewStaticCredentials(cfg.AWS.AccessKey, cfg.AWS.AccessSecret, ""),
	})

	// Call the NewAwsS3 function and assign the return value to a variable
	send := storage.NewAwsS3(sess)

	// Upload the file to S3.
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithError(err)
	}

	contentType := mime.TypeByExtension(filepath.Ext(header.Filename))
	data := send.Put(ctx, "lars-storage", header.Filename, fileBytes, false, contentType)
	if data != nil {
		logger.WarnWithContext(ctx, fmt.Sprintf("validation error %v", data), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithError(data)
	}
	var filepath = "https://" + "lars-storage" + "." + "is3.cloudhost.id/" + header.Filename
	fmt.Println("File successfully uploaded to S3! :", filepath)

	logger.InfoWithContext(ctx, fmt.Sprintf("success store data to database"), lf...)
	return *appctx.NewResponse().
		WithMsgKey(consts.RespSuccess)
}
