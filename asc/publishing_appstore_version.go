package asc

import (
	"context"
)

type appStoreVersionReleaseRequestCreateRequest struct {
	Type          string        `json:"type"`
	RelationShips relationships `json:"relationships"`
}

type relationships struct {
	Version appStoreVersion `json:"appStoreVersion"`
}
type appStoreVersion struct {
	Data data `json:"data"`
}
type data struct {
	VersionID   string `json:"id"`
	VersionType string `json:"type"`
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
	versiondata := data{
		VersionID:   versionID,
		VersionType: "appStoreVersions",
	}
	appStoreVersion := appStoreVersion{
		Data: versiondata,
	}

	relationshipsData := relationships{
		Version: appStoreVersion,
	}
	req := appStoreVersionReleaseRequestCreateRequest{
		Type:          "appStoreVersionReleaseRequests",
		RelationShips: relationshipsData,
	}

	res := new(AppStoreVersionReleaseRequestResponse)
	resp, err := s.client.post(ctx, "appStoreVersionReleaseRequests", newRequestBody(req), res)

	return res, resp, err
}
