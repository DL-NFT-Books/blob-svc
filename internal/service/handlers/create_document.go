package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/responses"
)

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	document, header, err := requests.NewCreateDocumentRequest(r)
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
	key := uuid.New().String() + "." + ext

	err = helpers.UploadFile(document, key, awsConfig)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewKeyResponse(key))
}
