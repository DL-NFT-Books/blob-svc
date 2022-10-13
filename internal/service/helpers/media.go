package helpers

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strings"
)

func CheckBannerMimeType(mtype string, r *http.Request) (string, error) {
	for _, el := range MimeTypes(r).AllowedBannerMimeTypes {
		if el == mtype {
			return strings.Split(mtype, "/")[1], nil
		}
	}
	return "", errors.New("invalid banner extension")
}

func CheckDocumentMimeType(mtype string, r *http.Request) (string, error) {
	for _, el := range MimeTypes(r).AllowedFileMimeTypes {
		if el == mtype {
			return strings.Split(mtype, "/")[1], nil
		}
	}
	return "", errors.New("invalid file extension")
}

func IsDocument(ext string, r *http.Request) bool {
	for _, el := range MimeTypes(r).AllowedFileMimeTypes {
		if strings.Split(el, "/")[1] == ext {
			return true
		}
	}
	return false
}
