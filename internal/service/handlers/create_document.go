package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/responses"
)

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	key, document, header, err := requests.NewCreateDocumentRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ext, err := helpers.CheckDocumentMimeType(header.Header.Get("Content-Type"), r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	awsConfig := helpers.AwsConfig(r)

	if key == "" {
		key = uuid.New().String()
	} else {
		// checking if key exists (only in case of custom keys, not uuid-generated)
		// to not overwrite the existing document

		exists, err := helpers.IsKeyExists(key+"."+ext, awsConfig)
		if err != nil || !exists {
			ape.RenderErr(w, problems.BadRequest(
				errors.New("Document with such key already exists or it cannot be checked"))...)
			return
		}
	}
	key += "." + ext

	err = helpers.UploadFile(document, key, awsConfig)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewKeyResponse(key))
}
