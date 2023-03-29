package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"github.com/dl-nft-books/blob-svc/internal/service/helpers"
	"github.com/dl-nft-books/blob-svc/internal/service/requests"
	"github.com/dl-nft-books/blob-svc/internal/service/responses"
	"github.com/dl-nft-books/blob-svc/resources"
)

func GetDocumentByKey(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDocumentByKeyRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	awsConfig := helpers.AwsConfig(r)

	exists, err := helpers.IsKeyExists(req.Key, awsConfig)
	if err != nil || !exists {
		helpers.Log(r).WithError(err).Debug("failed to check key existence or key was not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	url, err := helpers.GetUrl(req.Key, awsConfig)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewLinkResponse(url, resources.NewKeyInt64(1, resources.DOCUMENTS)))
}
