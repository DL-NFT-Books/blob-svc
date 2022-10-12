/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DocumentResponse struct {
	Key
	Attributes DocumentResponseAttributes `json:"attributes"`
}
type DocumentResponseResponse struct {
	Data     DocumentResponse `json:"data"`
	Included Included         `json:"included"`
}

type DocumentResponseListResponse struct {
	Data     []DocumentResponse `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustDocumentResponse - returns DocumentResponse from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDocumentResponse(key Key) *DocumentResponse {
	var documentResponse DocumentResponse
	if c.tryFindEntry(key, &documentResponse) {
		return &documentResponse
	}
	return nil
}
