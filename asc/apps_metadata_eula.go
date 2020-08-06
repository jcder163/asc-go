package asc

import (
	"fmt"
)

// EndUserLicenseAgreement defines model for EndUserLicenseAgreement.
type EndUserLicenseAgreement struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		Territories *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"territories,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// EndUserLicenseAgreementCreateRequest defines model for EndUserLicenseAgreementCreateRequest.
type EndUserLicenseAgreementCreateRequest struct {
	Attributes    EndUserLicenseAgreementCreateRequestAttributes    `json:"attributes"`
	Relationships EndUserLicenseAgreementCreateRequestRelationships `json:"relationships"`
	Type          string                                            `json:"type"`
}

// EndUserLicenseAgreementCreateRequestAttributes are attributes for EndUserLicenseAgreementCreateRequest
type EndUserLicenseAgreementCreateRequestAttributes struct {
	AgreementText string `json:"agreementText"`
}

// EndUserLicenseAgreementCreateRequestRelationships are relationships for EndUserLicenseAgreementCreateRequest
type EndUserLicenseAgreementCreateRequestRelationships struct {
	App struct {
		Data RelationshipsData `json:"data"`
	} `json:"app"`
	Territories struct {
		Data []RelationshipsData `json:"data"`
	} `json:"territories"`
}

// EndUserLicenseAgreementUpdateRequest defines model for EndUserLicenseAgreementUpdateRequest.
type EndUserLicenseAgreementUpdateRequest struct {
	Attributes    *EndUserLicenseAgreementUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                                             `json:"id"`
	Relationships *EndUserLicenseAgreementUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                             `json:"type"`
}

// EndUserLicenseAgreementUpdateRequestAttributes are attributes for EndUserLicenseAgreementUpdateRequest
type EndUserLicenseAgreementUpdateRequestAttributes struct {
	AgreementText *string `json:"agreementText,omitempty"`
}

// EndUserLicenseAgreementUpdateRequestRelationships are relationships for EndUserLicenseAgreementUpdateRequest
type EndUserLicenseAgreementUpdateRequestRelationships struct {
	Territories *struct {
		Data *[]RelationshipsData `json:"data,omitempty"`
	} `json:"territories,omitempty"`
}

// EndUserLicenseAgreementResponse defines model for EndUserLicenseAgreementResponse.
type EndUserLicenseAgreementResponse struct {
	Data     EndUserLicenseAgreement `json:"data"`
	Included *[]Territory            `json:"included,omitempty"`
	Links    DocumentLinks           `json:"links"`
}

// GetEULAQuery are query options for GetEULA
type GetEULAQuery struct {
	FieldsEndUserLicenseAgreements []string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsTerritories              []string `url:"fields[territories],omitempty"`
	Include                        []string `url:"include,omitempty"`
	LimitTerritories               int      `url:"limit[territories],omitempty"`
}

// GetEULAForAppQuery are query options for GetEULAForApp
type GetEULAForAppQuery struct {
	FieldsEndUserLicenseAgreements []string `url:"fields[endUserLicenseAgreements],omitempty"`
}

// CreateEULA adds a custom end user license agreement (EULA) to an app and configure the territories to which it applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_end_user_license_agreement
func (s *AppsService) CreateEULA(body *EndUserLicenseAgreementCreateRequest) (*EndUserLicenseAgreementResponse, *Response, error) {
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.post("endUserLicenseAgreements", body, res)
	return res, resp, err
}

// UpdateEULA updates the text or territories for your custom end user license agreement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_end_user_license_agreement
func (s *AppsService) UpdateEULA(id string, body *EndUserLicenseAgreementUpdateRequest) (*EndUserLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// DeleteEULA deletes the custom end user license agreement that is associated with an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_end_user_license_agreement
func (s *AppsService) DeleteEULA(id string) (*Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	return s.client.delete(url, nil)
}

// GetEULA gets the custom end user license agreement associated with an app, and the territories it applies to.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_end_user_license_agreement_information
func (s *AppsService) GetEULA(id string, params *GetEULAQuery) (*EndUserLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetEULAForApp gets the custom end user license agreement (EULA) for a specific app and the territories where the agreement applies.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_end_user_license_agreement_information_of_an_app
func (s *AppsService) GetEULAForApp(id string, params *GetEULAForAppQuery) (*EndUserLicenseAgreementResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/endUserLicenseAgreement", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}