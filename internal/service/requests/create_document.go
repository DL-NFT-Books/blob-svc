package requests

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var keyRegexp = regexp.MustCompile("^[^<>:;(),.?\"*|/]+$")

func NewCreateDocumentRequest(r *http.Request) (string, multipart.File, *multipart.FileHeader, error) {
	err := r.ParseMultipartForm(1 << 32)
	if err != nil {
		fmt.Println(err)
		return "", nil, nil, errors.Wrap(err, "failed to parse document")
	}

	key := r.FormValue("Key")
	if key != "" {

		err = validation.Errors{
			"key": validation.Validate(key, validation.Required, validation.Match(keyRegexp)),
		}.Filter()

		if err != nil {
			return "", nil, nil, errors.Wrap(err, "failed to parse key")
		}
	}

	f, h, err := r.FormFile("Document")
	return key, f, h, err
}
