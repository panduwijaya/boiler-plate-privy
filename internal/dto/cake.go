package dto

import (
	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/internal/entity"
	"cake-store/cake-store/internal/presentations"
	"cake-store/cake-store/pkg/storage"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var cfg = appctx.NewConfig()

func CakeToResponse(src entity.Cake) presentations.CakeDetail {
	x := presentations.CakeDetail{
		ID:          src.ID,
		Title:       src.Title,
		Description: src.Description,
		Rating:      src.Rating,
		Image:       src.Image,
		CreatedAt:   src.CreatedAt.String(),
		UpdatedAt:   src.UpdatedAt.String(),
	}

	if !src.CreatedAt.IsZero() {
		x.CreatedAt = src.CreatedAt.Format(consts.LayoutDateTimeFormat)
	}

	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String("https://is3.cloudhost.id"),
		Region:      aws.String(cfg.AWS.Region),
		Credentials: credentials.NewStaticCredentials(cfg.AWS.AccessKey, cfg.AWS.AccessSecret, ""),
	})

	// Call the NewAwsS3 function and assign the return value to a variable
	get_aws := storage.NewAwsS3(sess)
	data, err := get_aws.GetUrl(context.Background(), "lars-storage", src.Image)
	if err != nil {
		// handle error
	}
	x.Image = data

	//if !src.UpdatedAt.IsZero() {
	//	x.UpdatedAt = src.UpdatedAt.Format(consts.LayoutDateTimeFormat)
	//}

	//if !src.DeletedAt.IsZero() {
	//	x.DeletedAt = src.DeletedAt.Format(consts.LayoutDateTimeFormat)
	//}

	return x
}

func CakesToResponse(inputs []entity.Cake) []presentations.CakeDetail {
	var (
		result = []presentations.CakeDetail{}
	)

	for i := 0; i < len(inputs); i++ {
		result = append(result, CakeToResponse(inputs[i]))
	}

	return result
}
