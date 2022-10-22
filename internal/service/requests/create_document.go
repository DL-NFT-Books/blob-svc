package requests

import (
	"mime/multipart"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewCreateDocumentRequest(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	err := r.ParseMultipartForm(1 << 32)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse document")
	}

	return r.FormFile("Document")
}
