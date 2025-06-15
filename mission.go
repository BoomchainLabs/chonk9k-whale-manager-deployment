// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package earnapp

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/BoomchainLabs/release-please/internal/apijson"
	"github.com/BoomchainLabs/release-please/internal/apiquery"
	"github.com/BoomchainLabs/release-please/internal/requestconfig"
	"github.com/BoomchainLabs/release-please/option"
	"github.com/BoomchainLabs/release-please/packages/param"
	"github.com/BoomchainLabs/release-please/packages/respjson"
)

// MissionService contains methods and other services that help with interacting
// with the earn-app API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMissionService] method instead.
type MissionService struct {
	Options []option.RequestOption
}

// NewMissionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMissionService(opts ...option.RequestOption) (r MissionService) {
	r = MissionService{}
	r.Options = opts
	return
}

// Retrieve detailed information about a specific mission
func (r *MissionService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *Mission, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("missions/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the list of all available missions
func (r *MissionService) List(ctx context.Context, query MissionListParams, opts ...option.RequestOption) (res *[]Mission, err error) {
	opts = append(r.Options[:], opts...)
	path := "missions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Mission struct {
	// Unique mission identifier
	ID int64 `json:"id"`
	// Mission category
	Category string `json:"category"`
	// Detailed mission description
	Description string `json:"description"`
	// Mission end timestamp
	EndDate time.Time `json:"endDate" format:"date-time"`
	// Points rewarded for completing the mission
	Points int64 `json:"points"`
	// List of requirements to complete the mission
	Requirements []MissionRequirement `json:"requirements"`
	// Mission start timestamp
	StartDate time.Time `json:"startDate" format:"date-time"`
	// Current mission status
	//
	// Any of "active", "inactive", "completed".
	Status MissionStatus `json:"status"`
	// Mission title
	Title string `json:"title"`
	// Token amount rewarded for completing the mission
	TokenReward float64 `json:"tokenReward"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Category     respjson.Field
		Description  respjson.Field
		EndDate      respjson.Field
		Points       respjson.Field
		Requirements respjson.Field
		StartDate    respjson.Field
		Status       respjson.Field
		Title        respjson.Field
		TokenReward  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Mission) RawJSON() string { return r.JSON.raw }
func (r *Mission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MissionRequirement struct {
	// Unique requirement identifier
	ID int64 `json:"id"`
	// Requirement description
	Description string `json:"description"`
	// Target value to reach
	Target int64 `json:"target"`
	// Type of requirement
	//
	// Any of "transaction", "social", "quiz", "game", "other".
	Type string `json:"type"`
	// How the requirement is verified
	//
	// Any of "automatic", "manual".
	VerificationMethod string `json:"verificationMethod"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Description        respjson.Field
		Target             respjson.Field
		Type               respjson.Field
		VerificationMethod respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MissionRequirement) RawJSON() string { return r.JSON.raw }
func (r *MissionRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current mission status
type MissionStatus string

const (
	MissionStatusActive    MissionStatus = "active"
	MissionStatusInactive  MissionStatus = "inactive"
	MissionStatusCompleted MissionStatus = "completed"
)

type MissionListParams struct {
	// Filter missions by category
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Filter missions by status
	//
	// Any of "active", "inactive", "completed".
	Status MissionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MissionListParams]'s query parameters as `url.Values`.
func (r MissionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter missions by status
type MissionListParamsStatus string

const (
	MissionListParamsStatusActive    MissionListParamsStatus = "active"
	MissionListParamsStatusInactive  MissionListParamsStatus = "inactive"
	MissionListParamsStatusCompleted MissionListParamsStatus = "completed"
)
