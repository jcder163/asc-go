package asc

import (
	"context"
	"fmt"
)

// AppStoreReviewDetailResponse defines model for AppStoreReviewDetailResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailresponse
type ReviewSubmissionResponse struct {
	Data     ReviewSubmission           `json:"data"`
	Included []AppStoreReviewAttachment `json:"included,omitempty"`
	Links    DocumentLinks              `json:"links"`
}

// AppStoreReviewDetail defines model for AppStoreReviewDetail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetail
type ReviewSubmission struct {
	Attributes    *ReviewSubmissionAttributes    `json:"attributes,omitempty"`
	ID            string                         `json:"id"`
	Links         ResourceLinks                  `json:"links"`
	Relationships *ReviewSubmissionRelationships `json:"relationships,omitempty"`
	Type          string                         `json:"type"`
}

type ReviewSubmissionRelationships struct {
	App                      ReviewSubmissionRelationshipsApp                      `json:"app,omitempty"`
	AppStoreVersionForReview ReviewSubmissionRelationshipsAppStoreVersionForReview `json:"appStoreVersionForReview,omitempty"`
	Items                    ReviewSubmissionRelationshipsItems                    `json:"items,omitempty"`
}

type ReviewSubmissionRelationshipsApp struct {
	Data  RelationshipData  `json:"data,omitempty"`
	Links RelationshipLinks `json:"links,omitempty"`
}

type ReviewSubmissionRelationshipsAppStoreVersionForReview struct {
	Data  RelationshipData  `json:"data,omitempty"`
	Links RelationshipLinks `json:"links,omitempty"`
}

type ReviewSubmissionRelationshipsItems struct {
	Data  []RelationshipData `json:"data,omitempty"`
	Links RelationshipLinks  `json:"links,omitempty"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

type ReviewState string

const (
	ReviewStateReadyForReview   ReviewState = "READY_FOR_REVIEW"
	ReviewStateWaitingForReview ReviewState = "WAITING_FOR_REVIEW"
	ReviewStateInReview         ReviewState = "IN_REVIEW"
	ReviewStateUnResolvedIssues ReviewState = "UNRESOLVED_ISSUES"
	ReviewStateCanceling        ReviewState = "CANCELING"
	ReviewStateCompleting       ReviewState = "COMPLETING"
	ReviewStateComplete         ReviewState = "COMPLETE"
)

type ReviewSubmissionAttributes struct {
	Platform      Platform
	State         ReviewState
	SubmittedDate *DateTime `json:"earliestReleaseDate,omitempty"`
}

// GetReviewDetailQuery are query options for GetReviewDetail
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_review_detail_information
type ReadReviewSubmissionInfomationQuery struct {
	FieldsReviewSubmissionItems    []string `url:"fields[reviewSubmissionItems],omitempty"`
	FieldsReviewSubmissions        []string `url:"fields[reviewSubmissions],omitempty"`
	Include                        []string `url:"include,omitempty"`
	LimitAppStoreReviewAttachments int      `url:"limit[items],omitempty"`
}

// UpdateReviewDetail update the app store review details, including the contact information, demo account, and notes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_review_submission_information
func (s *SubmissionService) ReadReviewSubmissionInfomation(ctx context.Context, id string, params *ReadReviewSubmissionInfomationQuery) (*ReviewSubmissionResponse, *Response, error) {
	url := fmt.Sprintf("reviewSubmissions/%s", id)
	res := new(ReviewSubmissionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// AppStoreReviewDetailResponse defines model for AppStoreReviewDetailResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstorereviewdetailresponse
type ListReviewSubmissionResponse struct {
	Data     []ReviewSubmission         `json:"data"`
	Included []AppStoreReviewAttachment `json:"included,omitempty"`
	Links    DocumentLinks              `json:"links"`
	Meta     PagingInformation          `json:"meta"`
}

type ListReviewSubmissionInfomationQuery struct {
	FieldsReviewSubmissionItems []string `url:"fields[reviewSubmissionItems],omitempty"`
	FieldsReviewSubmissions     []string `url:"fields[reviewSubmissionItems],omitempty"`
	FilterAPP                   []string `url:"filter[app],omitempty"`
	FilterPlatform              []string `url:"filter[platform],omitempty"`
	FilterState                 []string `url:"filter[state],omitempty"`
	Include                     []string `url:"include,omitempty"`
	Limit                       int      `url:"limit,omitempty"`
	LimitItems                  int      `url:"limit[items],omitempty"`
}

// UpdateReviewDetail update the app store review details, including the contact information, demo account, and notes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_review_submission_information
func (s *SubmissionService) ListReviewSubmissionInfomation(ctx context.Context, params *ListReviewSubmissionInfomationQuery) (*ListReviewSubmissionResponse, *Response, error) {
	res := new(ListReviewSubmissionResponse)
	resp, err := s.client.get(ctx, "reviewSubmissions", params, res)

	return res, resp, err
}
