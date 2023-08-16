package cloudinary

import (
	"context"
	"errors"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"mime/multipart"
	"os"
)

type Media interface {
	Upload(file multipart.FileHeader) (*string, *response.Errors)
	Delete(file string) (*string, *response.Errors)
}

type mediaUseCase struct {
}

func (useCase *mediaUseCase) Upload(file multipart.FileHeader) (*string, *response.Errors) {
	//Define cloudinary access
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUDNAME"))

	if err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()

	binary, err := file.Open()
	if err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}

	defer func(binary multipart.File) {
		err := binary.Close()
		if err != nil {
			return
		}
	}(binary)

	if binary != nil {
		uploadResult, err := cld.Upload.Upload(
			ctx,
			binary,
			uploader.UploadParams{PublicID: uuid.New().String()},
		)
		if err != nil {
			return nil, &response.Errors{
				Code: 500,
				Err:  err,
			}
		}
		return &uploadResult.SecureURL, nil
	}
	return nil, &response.Errors{
		Code: 500,
		Err:  errors.New("can't read binary file"),
	}
}

func (useCase *mediaUseCase) Delete(file string) (*string, *response.Errors) {
	//Define cloudinary access
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUDNAME"))

	if err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()

	fileName := utils.GetFileName(file)
	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: fileName})

	if err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &res.Result, nil
}

func NewMediaUseCase() Media {
	return &mediaUseCase{}
}
