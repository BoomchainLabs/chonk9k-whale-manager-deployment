// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package earnapp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/BoomchainLabs/release-please/internal/apijson"
	"github.com/BoomchainLabs/release-please/internal/requestconfig"
	"github.com/BoomchainLabs/release-please/option"
	"github.com/BoomchainLabs/release-please/packages/param"
	"github.com/BoomchainLabs/release-please/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the earn-app API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
	Stakes  UserStakeService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	r.Stakes = NewUserStakeService(opts...)
	return
}

// Register a new user in the system with their wallet address
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all missions for a specific user
func (r *UserService) ListMissions(ctx context.Context, userID int64, opts ...option.RequestOption) (res *[]UserListMissionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("users/%v/missions", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all referrals made by a specific user
func (r *UserService) ListReferrals(ctx context.Context, userID int64, opts ...option.RequestOption) (res *[]UserListReferralsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("users/%v/referrals", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve user details by their blockchain wallet address
func (r *UserService) GetByWallet(ctx context.Context, address string, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("users/wallet/%s", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type User struct {
	// Unique user identifier
	ID int64 `json:"id"`
	// User creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// User's email address
	Email string `json:"email" format:"email"`
	// User's current level
	Level int64 `json:"level"`
	// User's unique referral code
	ReferralCode string `json:"referralCode"`
	// Total points earned by the user
	TotalPoints int64 `json:"totalPoints"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// User's chosen username
	Username string `json:"username"`
	// User's blockchain wallet address
	WalletAddress string `json:"walletAddress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Email         respjson.Field
		Level         respjson.Field
		ReferralCode  respjson.Field
		TotalPoints   respjson.Field
		UpdatedAt     respjson.Field
		Username      respjson.Field
		WalletAddress respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListMissionsResponse struct {
	// Unique progress identifier
	ID int64 `json:"id"`
	// Whether the reward has been claimed
	Claimed bool `json:"claimed"`
	// Reward claim timestamp
	ClaimedAt time.Time `json:"claimedAt" format:"date-time"`
	// Whether the mission is completed
	Completed bool `json:"completed"`
	// Completion timestamp
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	Mission     Mission   `json:"mission"`
	// Mission identifier
	MissionID int64 `json:"missionId"`
	// Current progress percentage (0-100)
	Progress float64 `json:"progress"`
	// User identifier
	UserID int64 `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Claimed     respjson.Field
		ClaimedAt   respjson.Field
		Completed   respjson.Field
		CompletedAt respjson.Field
		Mission     respjson.Field
		MissionID   respjson.Field
		Progress    respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListMissionsResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListMissionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListReferralsResponse struct {
	// Unique referral identifier
	ID int64 `json:"id"`
	// Referral code used
	Code string `json:"code"`
	// Referral creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	Referee   User      `json:"referee"`
	// Referred user identifier
	RefereeID int64 `json:"refereeId"`
	// Referrer user identifier
	ReferrerID int64 `json:"referrerId"`
	// Reward amount given to referrer
	RewardAmount float64 `json:"rewardAmount"`
	// Reward timestamp
	RewardedAt time.Time `json:"rewardedAt" format:"date-time"`
	// Referral status
	//
	// Any of "pending", "rewarded", "expired".
	Status UserListReferralsResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Code         respjson.Field
		CreatedAt    respjson.Field
		Referee      respjson.Field
		RefereeID    respjson.Field
		ReferrerID   respjson.Field
		RewardAmount respjson.Field
		RewardedAt   respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListReferralsResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListReferralsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Referral status
type UserListReferralsResponseStatus string

const (
	UserListReferralsResponseStatusPending  UserListReferralsResponseStatus = "pending"
	UserListReferralsResponseStatusRewarded UserListReferralsResponseStatus = "rewarded"
	UserListReferralsResponseStatusExpired  UserListReferralsResponseStatus = "expired"
)

type UserNewParams struct {
	// User's blockchain wallet address
	WalletAddress string `json:"walletAddress,required"`
	// User's email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Referral code used during signup
	ReferralCode param.Opt[string] `json:"referralCode,omitzero"`
	// User's chosen username
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
