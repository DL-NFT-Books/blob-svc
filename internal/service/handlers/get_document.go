package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/responses"
	"gitlab.com/tokend/nft-books/blob-svc/resources"
)

func GetFileByKey(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDocumentByKeyRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	awsConfig := helpers.AwsConfig(r)

	exists, err := helpers.IsKeyExists(req.Key, awsConfig)
	if err != nil || !exists {
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
