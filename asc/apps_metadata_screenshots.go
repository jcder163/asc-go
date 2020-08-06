package asc

import (
	"fmt"
)

// AppScreenshot defines model for AppScreenshot.
type AppScreenshot struct {
	Attributes *struct {
		AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
		AssetToken         *string             `json:"assetToken,omitempty"`
		AssetType          *string             `json:"assetType,omitempty"`
		FileName           *string             `json:"fileName,omitempty"`
		FileSize           *int64              `json:"fileSize,omitempty"`
		ImageAsset         *ImageAsset         `json:"imageAsset,omitempty"`
		SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
		UploadOperations   *UploadOperations   `json:"uploadOperations,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppScreenshotSet *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appScreenshotSet,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppScreenshotCreateRequest defines model for AppScreenshotCreateRequest.
type AppScreenshotCreateRequest struct {
	Attributes    AppScreenshotCreateRequestAttributes    `json:"attributes"`
	Relationships AppScreenshotCreateRequestRelationships `json:"relationships"`
	Type          string                                  `json:"type"`
}

// AppScreenshotCreateRequestAttributes are attributes for AppScreenshotCreateRequest
type AppScreenshotCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
}

// AppScreenshotCreateRequestRelationships are relationships for AppScreenshotCreateRequest
type AppScreenshotCreateRequestRelationships struct {
	AppScreenshotSet struct {
		Data RelationshipsData `json:"data"`
	} `json:"appScreenshotSet"`
}

// AppScreenshotUpdateRequest defines model for AppScreenshotUpdateRequest.
type AppScreenshotUpdateRequest struct {
	Attributes *AppScreenshotUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                `json:"id"`
	Type       string                                `json:"type"`
}

// AppScreenshotUpdateRequestAttributes are attributes for AppScreenshotUpdateRequest
type AppScreenshotUpdateRequestAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	Uploaded           *bool   `json:"uploaded,omitempty"`
}

// AppScreenshotResponse defines model for AppScreenshotResponse.
type AppScreenshotResponse struct {
	Data  AppScreenshot `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppScreenshotsResponse defines model for AppScreenshotsResponse.
type AppScreenshotsResponse struct {
	Data  []AppScreenshot    `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// GetAppScreenshotQuery are query options for GetAppScreenshot
type GetAppScreenshotQuery struct {
	FieldsAppScreenshots []string `url:"fields[appScreenshots],omitempty"`
	Include              []string `url:"include,omitempty"`
}

// GetAppScreenshot gets information about an app screenshot and its upload and processing status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_screenshot_information
func (s *AppsService) GetAppScreenshot(id string, params *GetAppScreenshotQuery) (*AppScreenshotResponse, *Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// CreateAppScreenshot adds a new screenshot to a screenshot set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_screenshot
func (s *AppsService) CreateAppScreenshot(body *AppScreenshotCreateRequest) (*AppScreenshotResponse, *Response, error) {
	res := new(AppScreenshotResponse)
	resp, err := s.client.post("appScreenshots", body, res)
	return res, resp, err
}

// CommitAppScreenshot commits an app screenshot after uploading it.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_screenshot
func (s *AppsService) CommitAppScreenshot(id string, body *AppScreenshotUpdateRequest) (*AppScreenshotResponse, *Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// DeleteAppScreenshot deletes an app screenshot that is associated with a screenshot set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_screenshot
func (s *AppsService) DeleteAppScreenshot(id string) (*Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	return s.client.delete(url, nil)
}