package handlers

import (
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/responses"
	"gitlab.com/tokend/nft-books/blob-svc/resources"
)

func GetFileByKey(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetFileByKeyRequest(r)
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

	var key resources.Key
	ext := strings.Split(req.Key, ".")[1]
	if helpers.IsDocument(ext, r) {
		key = resources.NewKeyInt64(1, resources.DOCUMENTS)
	} else {
		key = resources.NewKeyInt64(1, resources.IMAGES)
	}

	ape.Render(w, responses.NewLinkResponse(url, key))
}
