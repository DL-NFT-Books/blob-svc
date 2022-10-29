package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GetDocumentByKeyRequest struct {
	Key string
}

func NewGetDocumentByKeyRequest(r *http.Request) (GetDocumentByKeyRequest, error) {
	request := GetDocumentByKeyRequest{}

	key := chi.URLParam(r, "key")
	if len(key) == 0 {
		return GetDocumentByKeyRequest{}, errors.New("failed to retrieve key")
	}

	request.Key = key
	return request, nil
}
