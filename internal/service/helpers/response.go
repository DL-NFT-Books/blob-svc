package helpers

import "gitlab.com/tokend/nft-books/blob-svc/resources"

func NewKeyResponse(key string) resources.KeyResponseResponse {
	return resources.KeyResponseResponse{
		Data: resources.KeyResponse{
			Key: resources.NewKeyInt64(0, resources.S3KEY),
			Attributes: resources.KeyResponseAttributes{
				S3Key: key,
			},
		},
	}
}
