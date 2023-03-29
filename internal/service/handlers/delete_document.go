package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"github.com/dl-nft-books/blob-svc/internal/service/helpers"
	"github.com/dl-nft-books/blob-svc/internal/service/requests"
)

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDocumentByKeyRequest(r) // As only key is required
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	awsConfig := helpers.AwsConfig(r)

	// Make sure the key exists as the DeleteFile method will not render the KeyNotFound error
	exists, err := helpers.IsKeyExists(req.Key, awsConfig)
	if err != nil || !exists {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	err = helpers.DeleteFile(req.Key, awsConfig)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
