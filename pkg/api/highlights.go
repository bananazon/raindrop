package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/bananazon/raindrop/pkg/data"
)

func (ac *APIClient) ListHighlights(page int) (data.ListHighlightsResult, error) {
	var (
		err                  error
		listHighlightsResult data.ListHighlightsResult
		listUrl              url.URL
		queryMap             map[string]string
		response             APIResponse
	)

	queryMap = map[string]string{"perpage": strconv.Itoa(PageSize), "page": strconv.Itoa(page)}
	listUrl = url.URL{Scheme: "https", Host: apiBase, Path: fmt.Sprintf("rest/%s/highlights", apiVersion), RawQuery: MapToQueryString(queryMap)}
	response = ac.Request(APIRequest{Method: "GET", URL: listUrl})

	if !response.Success {
		return listHighlightsResult, response.Error
	}

	err = json.Unmarshal(response.Body, &listHighlightsResult)
	if err != nil {
		return listHighlightsResult, err
	}

	if !listHighlightsResult.Result {
		return listHighlightsResult, fmt.Errorf("list hightlights returned false: %s", listHighlightsResult.ErrorMessage)
	}

	return listHighlightsResult, nil
}

func (ac *APIClient) RemoveHighlight(bookmarkId int64, payload data.RemoveHighlightsPayload) (data.RemoveHighlightResult, error) {
	var (
		err                   error
		removeHighlightResult data.RemoveHighlightResult
		removeUrl             url.URL
		response              APIResponse
	)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return removeHighlightResult, err
	}

	removeUrl = url.URL{Scheme: "https", Host: apiBase, Path: fmt.Sprintf("rest/%s/raindrop/%d", apiVersion, bookmarkId)}
	response = ac.Request(APIRequest{Method: "PUT", URL: removeUrl, Body: string(jsonData)})
	if !response.Success {
		return removeHighlightResult, response.Error
	}

	err = json.Unmarshal(response.Body, &removeHighlightResult)
	if err != nil {
		return removeHighlightResult, err
	}

	if !removeHighlightResult.Result {
		return removeHighlightResult, fmt.Errorf("remove highlight returned false: %s", removeHighlightResult.ErrorMessage)
	}

	return removeHighlightResult, nil
}
