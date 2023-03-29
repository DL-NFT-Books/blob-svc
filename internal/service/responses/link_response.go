package responses

import "github.com/dl-nft-books/blob-svc/resources"

func NewLinkResponse(url string, key resources.Key) resources.LinkResponse {
	return resources.LinkResponse{
		Data: resources.Link{
			Key: key,
			Attributes: resources.LinkAttributes{
				Url: url,
			},
		},
	}
}
