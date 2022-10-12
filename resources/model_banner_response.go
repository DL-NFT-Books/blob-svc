/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BannerResponse struct {
	Key
	Attributes BannerResponseAttributes `json:"attributes"`
}
type BannerResponseResponse struct {
	Data     BannerResponse `json:"data"`
	Included Included       `json:"included"`
}

type BannerResponseListResponse struct {
	Data     []BannerResponse `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustBannerResponse - returns BannerResponse from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBannerResponse(key Key) *BannerResponse {
	var bannerResponse BannerResponse
	if c.tryFindEntry(key, &bannerResponse) {
		return &bannerResponse
	}
	return nil
}
