package utils

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

var cld *cloudinary.Cloudinary

func InitCloudinary(cloud_name, api_key, api_secret string) {
	var err error
	cld, err = cloudinary.NewFromParams(cloud_name, api_key, api_secret)
	if err != nil {
		log.Fatal(err)
	}
}

func UploadBase64Image(b64 string) (string, error) {
	res, err := cld.Upload.Upload(context.Background(), b64, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return res.SecureURL, err
}
