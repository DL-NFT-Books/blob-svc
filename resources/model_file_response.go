/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type FileResponse struct {
	Key
	Attributes FileResponseAttributes `json:"attributes"`
}
type FileResponseResponse struct {
	Data     FileResponse `json:"data"`
	Included Included     `json:"included"`
}

type FileResponseListResponse struct {
	Data     []FileResponse `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustFileResponse - returns FileResponse from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustFileResponse(key Key) *FileResponse {
	var fileResponse FileResponse
	if c.tryFindEntry(key, &fileResponse) {
		return &fileResponse
	}
	return nil
}
