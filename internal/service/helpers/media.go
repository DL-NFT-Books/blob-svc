package helpers

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strings"
)

func CheckBannerMimeType(ext string, r *http.Request) (string, error) {
	for _, el := range MimeTypes(r).AllowedBannerMimeTypes {
		if el == ext {
			return strings.Split(ext, "/")[1], nil
		}
	}
	return "", errors.New("invalid banner extension")
}

func CheckDocumentMimeType(ext string, r *http.Request) (string, error) {
	for _, el := range MimeTypes(r).AllowedFileMimeTypes {
		if el == ext {
			return strings.Split(ext, "/")[1], nil
		}
	}
	return "", errors.New("invalid file extension")
}
