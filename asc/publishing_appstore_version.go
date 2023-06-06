package asc

import (
	"context"
)

type appStoreVersionReleaseRequestCreateRequest struct {
	Type        string `json:"data.type"`
	VersionID   string `json:"data.relationships.appStoreVersion.data.id"`
	VersionType string `json:"data.relationships.appStoreVersion.data.type"`
}

type AppStoreVersionReleaseRequest struct {
	ID    string        `json:"ID"`
	Type  string        `json:"type"`
	Links ResourceLinks `json:"links"`
}

type AppStoreVersionReleaseRequestResponse struct {
	Data  AppStoreVersionReleaseRequest `json:"data"`
	Links DocumentLinks                 `json:"lins"`
}

func (s *PublishingService) PublishingAppVersion(ctx context.Context, versionID string) (*AppStoreVersionReleaseRequestResponse, *Response, error) {
	req := appStoreVersionReleaseRequestCreateRequest{
		Type:        "appStoreVersionReleaseRequests",
		VersionID:   versionID,
		VersionType: "appStoreVersions",
	}

	res := new(AppStoreVersionReleaseRequestResponse)
	resp, err := s.client.post(ctx, "appStoreVersionReleaseRequests", newRequestBody(req), res)

	return res, resp, err
}
