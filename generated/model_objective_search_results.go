/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// The results of the Objective search query.
type ObjectiveSearchResults struct {
	// The total number of matches for the search query. The first 1000 matches can be paged through via the API.
	Total int64 `json:"total"`
	// A list of search results.
	Data []ObjectiveSearchResult `json:"data"`
	// The URL path and query string for the next page of search results.
	Next string `json:"next"`
}
