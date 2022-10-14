package handlers

import (
	"fmt"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/requests"
	"net/http"
)

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetFileByKeyRequest(r) // As only id is required
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
		fmt.Println(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
