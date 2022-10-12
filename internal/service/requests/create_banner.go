package requests

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"mime/multipart"
	"net/http"
)

func NewCreateBannerRequest(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	err := r.ParseMultipartForm(1 << 32)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse document")
	}

	return r.FormFile("Banner")
}
