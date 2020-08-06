package asc

import (
	"fmt"
	"time"
)

// UserInvitation defines model for UserInvitation.
type UserInvitation struct {
	Attributes *struct {
		AllAppsVisible      *bool       `json:"allAppsVisible,omitempty"`
		Email               *Email      `json:"email,omitempty"`
		ExpirationDate      *time.Time  `json:"expirationDate,omitempty"`
		FirstName           *string     `json:"firstName,omitempty"`
		LastName            *string     `json:"lastName,omitempty"`
		ProvisioningAllowed *bool       `json:"provisioningAllowed,omitempty"`
		Roles               *[]UserRole `json:"roles,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		VisibleApps *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"visibleApps,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// UserInvitationCreateRequest defines model for UserInvitationCreateRequest.
type UserInvitationCreateRequest struct {
	Attributes    UserInvitationCreateRequestAttributes     `json:"attributes"`
	Relationships *UserInvitationCreateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                    `json:"type"`
}

// UserInvitationCreateRequestAttributes are attributes for UserInvitationCreateRequest
type UserInvitationCreateRequestAttributes struct {
	AllAppsVisible      *bool      `json:"allAppsVisible,omitempty"`
	Email               Email      `json:"email"`
	FirstName           string     `json:"firstName"`
	LastName            string     `json:"lastName"`
	ProvisioningAllowed *bool      `json:"provisioningAllowed,omitempty"`
	Roles               []UserRole `json:"roles"`
}

// UserInvitationCreateRequestRelationships are relationships for UserInvitationCreateRequest
type UserInvitationCreateRequestRelationships struct {
	VisibleApps *struct {
		Data *[]RelationshipsData `json:"data,omitempty"`
	} `json:"visibleApps,omitempty"`
}

// UserInvitationResponse defines model for UserInvitationResponse.
type UserInvitationResponse struct {
	Data     UserInvitation `json:"data"`
	Included *[]App         `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// UserInvitationsResponse defines model for UserInvitationsResponse.
type UserInvitationsResponse struct {
	Data     []UserInvitation   `json:"data"`
	Included *[]App             `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListInvitationsQuery is the query params structure for ListInvitations
type ListInvitationsQuery struct {
	FieldsApps            []string `url:"fields[apps],omitempty"`
	FieldsUserInvitations []string `url:"fields[userInvitations],omitempty"`
	FilterRoles           []string `url:"filter[roles],omitempty"`
	FilterEmail           []string `url:"filter[email],omitempty"`
	FilterVisibleApps     []string `url:"filter[visibleApps],omitempty"`
	Include               []string `url:"include,omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	LimitVisibleApps      int      `url:"limit[visibleApps],omitempty"`
	Sort                  []string `url:"sort,omitempty"`
	Cursor                string   `url:"cursor,omitempty"`
}

// GetInvitationQuery is the query params structure for GetInvitation
type GetInvitationQuery struct {
	FieldsApps            []string `url:"fields[apps],omitempty"`
	FieldsUserInvitations []string `url:"fields[userInvitations],omitempty"`
	Include               []string `url:"include,omitempty"`
	LimitVisibleApps      int      `url:"limit[visibleApps],omitempty"`
}

// ListInvitations gets a list of pending invitations to join your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_invited_users
func (s *UsersService) ListInvitations(params *ListInvitationsQuery) (*UserInvitationsResponse, *Response, error) {
	res := new(UserInvitationsResponse)
	resp, err := s.client.get("userInvitations", params, res)
	return res, resp, err
}

// GetInvitation gets information about a pending invitation to join your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_user_invitation_information
func (s *UsersService) GetInvitation(id string, params *GetInvitationQuery) (*UserInvitationResponse, *Response, error) {
	url := fmt.Sprintf("userInvitations/%s", id)
	res := new(UserInvitationResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// CreateInvitation invites a user with assigned user roles to join your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/invite_a_user
func (s *UsersService) CreateInvitation(body *UserInvitationCreateRequest) (*UserInvitationResponse, *Response, error) {
	res := new(UserInvitationResponse)
	resp, err := s.client.post("userInvitations", body, res)
	return res, resp, err
}

// CancelInvitation cancels a pending invitation for a user to join your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/cancel_a_user_invitation
func (s *UsersService) CancelInvitation(id string) (*Response, error) {
	url := fmt.Sprintf("userInvitations/%s", id)
	return s.client.delete(url, nil)
}

// ListVisibleAppsForInvitation gets a list of apps that will be visible to a user with a pending invitation.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_apps_visible_to_an_invited_user
func (s *UsersService) ListVisibleAppsForInvitation(id string, params ListVisibleAppsQuery) (*AppsResponse, *Response, error) {
	url := fmt.Sprintf("userInvitations/%s/visibleApps", id)
	res := new(AppsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}