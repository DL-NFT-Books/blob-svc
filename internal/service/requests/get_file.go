package requests

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type GetFileByKeyRequest struct {
	Key string
}

func NewGetFileByKeyRequest(r *http.Request) (GetFileByKeyRequest, error) {
	request := GetFileByKeyRequest{}

	key := chi.URLParam(r, "key")
	if len(key) == 0 {
		return GetFileByKeyRequest{}, errors.New("failed to retrieve key")
	}

	request.Key = key
	return request, nil
}
