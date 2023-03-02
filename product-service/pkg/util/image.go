package util

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func ParseImage(file *multipart.FileHeader) ([]byte, string, string, error) {
	// get the image data
	image, err := file.Open()
	if err != nil {
		return []byte{}, "", "", err
	}
	defer image.Close()

	// get the image type (extension)
	ext := filepath.Ext(file.Filename)
	imageType := strings.TrimPrefix(ext, ".")

	// get the image name
	imageName := strings.TrimSuffix(file.Filename, ext)

	imageData, err := ioutil.ReadAll(image)
	if err != nil {
		return []byte{}, "", "", err
	}

	return imageData, imageName, imageType, nil
}
